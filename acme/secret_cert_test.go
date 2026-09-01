package acme

import (
	"context"
	"testing"

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
