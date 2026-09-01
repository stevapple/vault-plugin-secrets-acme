package acme

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

// Renewal defers to Vault: the backend returns the lease as it stands and core
// works out the new term, bounded by the MaxTTL that was set from the
// certificate's NotAfter at issuance. Nothing the backend puts in TTL here can
// override a requested increment, so it must not try.
func TestCertRenewReturnsTheLeaseUnchanged(t *testing.T) {
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	b, err := Factory("test")(context.Background(), config)
	require.NoError(t, err)

	secret := &logical.Secret{
		InternalData: map[string]interface{}{"secret_type": secretCertType},
		LeaseOptions: logical.LeaseOptions{
			TTL:       48 * time.Hour,
			MaxTTL:    90 * 24 * time.Hour,
			Renewable: true,
			Increment: 24 * time.Hour,
		},
	}

	resp, err := b.HandleRequest(context.Background(), &logical.Request{
		Operation: logical.RenewOperation,
		Path:      "certs/lenstra.fr",
		Storage:   config.StorageView,
		Secret:    secret,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Secret)

	require.Equal(t, 48*time.Hour, resp.Secret.TTL, "the increment must not be folded into the TTL")
	require.Equal(t, 90*24*time.Hour, resp.Secret.MaxTTL, "the ceiling comes from the certificate and must survive")
}
