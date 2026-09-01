package acme

import (
	"context"
	"testing"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

// A lease can outlive the cache entry it was issued against: Cache.Read drops
// entries once the certificate passes cache_for_ratio of its lifetime, roles
// can set disable_cache, and the cache endpoint can clear everything. Revoking
// such a lease must not panic on the missing entry.
func TestRevokeWithoutCacheEntry(t *testing.T) {
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	b, err := Factory("test")(context.Background(), config)
	require.NoError(t, err)

	req := &logical.Request{
		Operation: logical.RevokeOperation,
		Path:      "certs/lenstra.fr",
		Storage:   config.StorageView,
		Secret: &logical.Secret{
			InternalData: map[string]interface{}{
				"secret_type": secretCertType,
				"cache_key":   cachePrefix + "0badc0de",
				"account":     "accounts/lenstra",
				"cert":        "",
			},
		},
	}

	resp, err := b.HandleRequest(context.Background(), req)
	require.NoError(t, err)
	require.Nil(t, resp)
}

func TestRevokeOnLeaseExpiry(t *testing.T) {
	tcases := []struct {
		Name     string
		Role     *role
		Lease    map[string]interface{}
		Expected bool
	}{
		{
			Name:     "role opted in",
			Role:     &role{Account: "lenstra", RevokeOnLeaseExpiry: true},
			Lease:    map[string]interface{}{"role": "roles/lenstra.fr"},
			Expected: true,
		},
		{
			Name:     "role did not opt in",
			Role:     &role{Account: "lenstra"},
			Lease:    map[string]interface{}{"role": "roles/lenstra.fr"},
			Expected: false,
		},
		{
			Name:     "role deleted while the lease was outstanding",
			Lease:    map[string]interface{}{"role": "roles/lenstra.fr"},
			Expected: false,
		},
		{
			// Leases issued before the setting existed.
			Name:     "lease carries no role",
			Role:     &role{Account: "lenstra", RevokeOnLeaseExpiry: true},
			Lease:    map[string]interface{}{},
			Expected: false,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.Name, func(t *testing.T) {
			config := logical.TestBackendConfig()
			config.StorageView = &logical.InmemStorage{}
			raw, err := Factory("test")(context.Background(), config)
			require.NoError(t, err)
			b, ok := raw.(backend)
			require.True(t, ok)

			if tc.Role != nil {
				require.NoError(t, tc.Role.save(context.Background(), config.StorageView, "roles/lenstra.fr"))
			}

			req := &logical.Request{
				Storage: config.StorageView,
				Secret:  &logical.Secret{InternalData: tc.Lease},
			}

			revoke, err := b.revokeOnLeaseExpiry(context.Background(), req)
			require.NoError(t, err)
			require.Equal(t, tc.Expected, revoke)
		})
	}
}

// getCacheKey hashes the role, so a field that does not change the certificate
// being issued must stay out of the JSON encoding — otherwise flipping it
// orphans every cache entry under the role and forces a re-issuance.
func TestCacheKeyIgnoresRevokeOnLeaseExpiry(t *testing.T) {
	data := &framework.FieldData{
		Raw:    map[string]interface{}{},
		Schema: map[string]*framework.FieldSchema{},
	}

	off, err := getCacheKey(&role{Account: "lenstra", AllowedDomains: []string{"lenstra.fr"}}, data)
	require.NoError(t, err)

	on, err := getCacheKey(&role{Account: "lenstra", AllowedDomains: []string{"lenstra.fr"}, RevokeOnLeaseExpiry: true}, data)
	require.NoError(t, err)

	require.Equal(t, off, on)
}

// The setting must survive the mapstructure round-trip through storage, which
// is separate from the JSON encoding the cache key uses.
func TestRoleRevokeOnLeaseExpiryRoundTrips(t *testing.T) {
	storage := &logical.InmemStorage{}
	ctx := context.Background()

	require.NoError(t, (&role{Account: "lenstra", RevokeOnLeaseExpiry: true}).save(ctx, storage, "roles/lenstra.fr"))

	r, err := getRole(ctx, storage, "roles/lenstra.fr")
	require.NoError(t, err)
	require.NotNil(t, r)
	require.True(t, r.RevokeOnLeaseExpiry)
}
