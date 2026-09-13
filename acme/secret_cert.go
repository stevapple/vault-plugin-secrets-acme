// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-acme/lego/v5/certcrypto"
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
			"domains": {
				Type: framework.TypeCommaStringSlice,
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

// roleForLease returns the role a lease was issued against, or nil when it
// cannot be established — the lease predates the role being recorded, or the
// role has since been deleted. Callers treat nil as "leave the certificate
// valid", the safe reading when nobody is left to state an intent.
//
// The role is re-read rather than snapshotted into the lease so that changing
// revoke_on_lease_expiry takes effect for leases already outstanding.
func (b *backend) roleForLease(ctx context.Context, req *logical.Request) (*role, error) {
	rolePath, ok := req.Secret.InternalData["role"].(string)
	if !ok || rolePath == "" {
		b.Logger().Debug("Lease carries no role, leaving the certificate valid")
		return nil, nil
	}

	r, err := getRole(ctx, req.Storage, rolePath)
	if err != nil {
		return nil, err
	}
	if r == nil {
		b.Logger().Debug("Role is gone, leaving the certificate valid", "role", rolePath)
	}

	return r, nil
}

func (b *backend) revokeLeaseCertificate(ctx context.Context, req *logical.Request, accountPath, cert string) error {
	// A lease lasts as long as its certificate, so by the time it ends the
	// certificate has often just expired. There is nothing left to revoke
	// then, and a CA that refuses to revoke an expired certificate would
	// otherwise keep this lease in Vault's revocation queue for good.
	if expired, err := certificateExpired(cert); err == nil && expired {
		b.Logger().Debug("Certificate has expired, nothing to revoke")
		return nil
	}

	a, err := getAccount(ctx, req.Storage, accountPath)
	if err != nil {
		return err
	}
	if a == nil {
		return fmt.Errorf("error while revoking certificate: user not found")
	}
	client, err := a.getClient()
	if err != nil {
		return fmt.Errorf("failed to get LEGO client: %w", err)
	}
	if err = client.Certificate.Revoke(ctx, []byte(cert)); err != nil {
		if alreadyRevoked(err) {
			// Through acme/revoke, another client, or the CA. The outcome is
			// the one the lease wanted, and failing here would leave Vault
			// retrying the lease revocation indefinitely.
			b.Logger().Debug("Certificate was already revoked, nothing left to do")
			return nil
		}
		return fmt.Errorf("failed to revoke cert: %v", err)
	}

	return nil
}

func (b *backend) certRevoke(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	b.cache.Lock()
	defer b.cache.Unlock()
	// Validate the lease before touching anything so that a malformed one
	// fails without leaving the cache entry half updated.
	cacheKey, ok := req.Secret.InternalData["cache_key"].(string)
	if !ok {
		return nil, errors.New("lease is missing its cache_key")
	}
	accountPath, ok := req.Secret.InternalData["account"].(string)
	if !ok {
		return nil, errors.New("lease is missing its account")
	}
	cert, ok := req.Secret.InternalData["cert"].(string)
	if !ok {
		return nil, errors.New("lease is missing its certificate")
	}

	ce, err := b.cache.Read(ctx, req.Storage, nil, cacheKey)
	if err != nil {
		return nil, err
	}
	if ce == nil {
		// No reference count to consult. A role with disable_cache never had
		// one: its leases hold a certificate of their own, so this lease going
		// away is the end of that certificate and the role's decision applies
		// directly.
		r, err := b.roleForLease(ctx, req)
		if err != nil {
			return nil, err
		}
		if r != nil && r.DisableCache && r.RevokeOnLeaseExpiry {
			return nil, b.revokeLeaseCertificate(ctx, req, accountPath, cert)
		}

		// Otherwise the entry was dropped once the certificate passed
		// cache_for_ratio of its lifetime, or the cache was cleared. Other
		// leases may well still hold this certificate and we can no longer
		// tell, so leave it valid.
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

		r, err := b.roleForLease(ctx, req)
		if err != nil {
			return nil, err
		}
		if r == nil || !r.RevokeOnLeaseExpiry {
			return nil, nil
		}

		if err = b.revokeLeaseCertificate(ctx, req, accountPath, cert); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// certificateExpired reports whether the leaf in a PEM bundle is past its
// NotAfter. An unparsable bundle is reported as an error so the caller can
// fall through to the provider, which will say what is wrong with it.
func certificateExpired(pemBundle string) (bool, error) {
	certs, err := certcrypto.ParsePEMBundle([]byte(pemBundle))
	if err != nil {
		return false, err
	}
	if len(certs) == 0 {
		return false, errors.New("no certificate in the PEM data")
	}
	return time.Now().After(certs[0].NotAfter), nil
}
