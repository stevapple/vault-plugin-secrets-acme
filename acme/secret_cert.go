// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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

func (b *backend) certRenew(_ context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	resp := &logical.Response{Secret: req.Secret}
	// I'm not really sure about this
	resp.Secret.TTL = resp.Secret.TTL + req.Secret.Increment
	return resp, nil
}

// revokeOnLeaseExpiry reports whether the role the lease was issued under
// wants the certificate revoked at the ACME provider now that its last lease
// is going away. The role is re-read rather than snapshotted into the lease so
// that turning the setting off takes effect for leases already outstanding.
//
// Leases issued before this setting existed carry no role, and a role can be
// deleted while its certificates are still in use. Both answer no: revocation
// is the destructive option, so anything short of an explicit opt-in leaves
// the certificate alone.
func (b *backend) revokeOnLeaseExpiry(ctx context.Context, req *logical.Request) (bool, error) {
	rolePath, ok := req.Secret.InternalData["role"].(string)
	if !ok || rolePath == "" {
		b.Logger().Debug("Lease carries no role, leaving the certificate valid")
		return false, nil
	}

	r, err := getRole(ctx, req.Storage, rolePath)
	if err != nil {
		return false, err
	}
	if r == nil {
		b.Logger().Debug("Role is gone, leaving the certificate valid", "role", rolePath)
		return false, nil
	}
	if !r.RevokeOnLeaseExpiry {
		b.Logger().Debug("Role does not enable revoke_on_lease_expiry, leaving the certificate valid", "role", rolePath)
		return false, nil
	}

	return true, nil
}

func (b *backend) certRevoke(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	b.cache.Lock()
	defer b.cache.Unlock()
	cacheKey := req.Secret.InternalData["cache_key"].(string)

	ce, err := b.cache.Read(ctx, req.Storage, nil, cacheKey)
	if err != nil {
		return nil, err
	}
	if ce == nil {
		// The entry is already gone: a later request found it stale and
		// dropped it, the role has disable_cache set so it was never written,
		// or the cache was cleared. Without the reference count we cannot tell
		// whether other leases still hold this certificate, so leave it valid
		// and let the lease go away.
		b.Logger().Debug("No cache entry for the revoked lease, nothing to do", "key", cacheKey)
		return nil, nil
	}

	ce.Users--
	if ce.Users > 0 {
		err = ce.Save(ctx, req.Storage, cacheKey)
		if err != nil {
			return nil, err
		}
	} else {
		// The last user asked for the lease to be terminated, so the cached
		// cert is of no further use to anyone.
		b.Logger().Debug("Removing cached cert", "key", cacheKey)
		err = b.cache.Delete(ctx, req.Storage, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("failed to remove cache entry: %v", err)
		}

		revoke, err := b.revokeOnLeaseExpiry(ctx, req)
		if err != nil {
			return nil, err
		}
		if !revoke {
			return nil, nil
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
