package acme

import (
	"context"
	"fmt"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

const secretCertType = "cert"

func secretCert(b *backend) *framework.Secret {
	return &framework.Secret{
		Type: secretCertType,
		Fields: map[string]*framework.FieldSchema{
			"domain": {
				Type: framework.TypeString,
			},
			"url": {
				Type: framework.TypeString,
			},
			"private_key": {
				Type: framework.TypeString,
			},
			"cert": {
				Type: framework.TypeString,
			},
			"issuer_cert": {
				Type: framework.TypeString,
			},
			"not_before": {
				Type: framework.TypeString,
			},
			"not_after": {
				Type: framework.TypeString,
			},
		},
		Renew:  b.certRenew,
		Revoke: b.certRevoke,
	}
}

// certRenew hands the lease back and lets Vault work out the new term. There
// is nothing for the backend to decide: the lease may be extended for as long
// as the certificate is valid, and MaxTTL was set from NotAfter when it was
// issued, so core will not extend one past the certificate it carries.
func (b *backend) certRenew(_ context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	return &logical.Response{Secret: req.Secret}, nil
}

func (b *backend) certRevoke(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	b.cache.Lock()
	defer b.cache.Unlock()
	cacheKey := req.Secret.InternalData["cache_key"].(string)

	ce, err := b.cache.Read(ctx, req.Storage, nil, cacheKey)
	if err != nil {
		return nil, err
	}

	ce.Users--
	if ce.Users > 0 {
		err = ce.Save(ctx, req.Storage, cacheKey)
		if err != nil {
			return nil, err
		}
	} else {
		// If the last user asked for the lease to be terminated we revoke the cert
		b.Logger().Debug("Removing cached cert", "key", cacheKey)
		err = b.cache.Delete(ctx, req.Storage, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("failed to remove cache entry: %v", err)
		}

		accountPath := req.Secret.InternalData["account"].(string)
		a, err := getAccount(ctx, req.Storage, accountPath)
		if err != nil {
			return nil, err
		}
		if a == nil {
			return nil, fmt.Errorf("error while revoking certificate: user not found")
		}
		client, err := a.getClient()
		if err != nil {
			return logical.ErrorResponse("Failed to get LEGO client."), err
		}
		cert := req.Secret.InternalData["cert"].(string)
		err = client.Certificate.Revoke([]byte(cert))
		if err != nil {
			return nil, fmt.Errorf("failed to revoke cert: %v", err)
		}
	}

	return nil, nil
}
