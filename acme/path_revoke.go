package acme

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-acme/lego/v5/certcrypto"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathRevoke(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: "revoke",
		Fields: map[string]*framework.FieldSchema{
			"account": {
				Type:        framework.TypeString,
				Required:    true,
				Description: "ACME account to revoke the certificate with.",
			},
			"certificate": {
				Type:        framework.TypeString,
				Required:    true,
				Description: "PEM encoded certificate to revoke.",
			},
			"reason": {
				Type:        framework.TypeInt,
				Default:     -1,
				Description: "RFC 5280 reason code to record at the ACME provider: 0 unspecified, 1 keyCompromise, 3 affiliationChanged, 4 superseded, 5 cessationOfOperation. Omitted by default.",
			},
		},
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.UpdateOperation: &framework.PathOperation{
				Callback: b.certRevokeNow,
			},
		},
	}
}

// certRevokeNow revokes a certificate on demand, independently of the leases
// carrying it. Roles that leave revoke_on_lease_expiry off have no other way
// to withdraw a certificate that has been compromised.
func (b *backend) certRevokeNow(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	accountName := data.Get("account").(string)
	if accountName == "" {
		return logical.ErrorResponse("account is required"), nil
	}
	certificate := data.Get("certificate").(string)
	if strings.TrimSpace(certificate) == "" {
		return logical.ErrorResponse("certificate is required"), nil
	}
	reason, err := revocationReason(data.Get("reason").(int))
	if err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	a, err := getAccount(ctx, req.Storage, "accounts/"+accountName)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return logical.ErrorResponse("This account does not exist"), nil
	}

	client, err := a.getClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get LEGO client: %w", err)
	}
	if err = client.Certificate.RevokeWithReason(ctx, []byte(certificate), reason); err != nil {
		return nil, fmt.Errorf("failed to revoke cert: %w", err)
	}

	// Drop it from the cache too, or the engine carries on handing out a
	// certificate it has just revoked.
	purged, err := b.purgeCachedCert(ctx, req.Storage, certificate)
	if err != nil {
		return nil, err
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"cache_entries_removed": purged,
		},
	}, nil
}

func (b *backend) purgeCachedCert(ctx context.Context, storage logical.Storage, certificate string) (int, error) {
	b.cache.Lock()
	defer b.cache.Unlock()

	keys, err := b.cache.List(ctx, storage)
	if err != nil {
		return 0, fmt.Errorf("failed to list the cache: %w", err)
	}

	want, err := leafDER(certificate)
	if err != nil {
		return 0, err
	}

	var purged int
	for _, key := range keys {
		// role nil: a plain lookup, with none of the staleness handling or
		// reference counting a certificate request would want.
		ce, err := b.cache.Read(ctx, storage, nil, cachePrefix+key)
		if err != nil {
			return purged, err
		}
		if ce == nil {
			continue
		}
		// Compare the leaf's DER rather than the PEM text: a cache entry holds
		// the whole bundle, and the caller may hand us either that or the leaf
		// on its own, with whatever line endings it picked up on the way.
		cached, err := leafDER(string(ce.Cert))
		if err != nil || !bytes.Equal(cached, want) {
			continue
		}
		if err = b.cache.Delete(ctx, storage, cachePrefix+key); err != nil {
			return purged, fmt.Errorf("failed to remove cache entry: %w", err)
		}
		purged++
	}

	return purged, nil
}

func leafDER(pemBundle string) ([]byte, error) {
	certs, err := certcrypto.ParsePEMBundle([]byte(pemBundle))
	if err != nil {
		return nil, fmt.Errorf("failed to parse the certificate: %w", err)
	}
	if len(certs) == 0 {
		return nil, errors.New("no certificate found in the PEM data")
	}

	return certs[0].Raw, nil
}

// revocationReason maps the request's reason to what lego sends: nil when
// none was given, so the provider records nothing, or the RFC 5280 CRLReason
// code. Code 7 is unassigned there and anything above 10 is not a code.
func revocationReason(code int) (*uint, error) {
	switch {
	case code < 0:
		return nil, nil
	case code == 7 || code > 10:
		return nil, fmt.Errorf("%d is not an RFC 5280 revocation reason", code)
	}
	reason := uint(code)
	return &reason, nil
}
