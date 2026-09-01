package acme

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

// certificateAtLifetime returns a self-signed certificate whose validity is
// placed so that the given share of its lifetime has already elapsed.
func certificateAtLifetime(t *testing.T, elapsed float64) []byte {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	lifetime := 90 * 24 * time.Hour
	notBefore := time.Now().Add(-time.Duration(elapsed * float64(lifetime)))
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "www.lenstra.fr"},
		NotBefore:    notBefore,
		NotAfter:     notBefore.Add(lifetime),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	require.NoError(t, err)

	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
}

// cache_for_ratio is the share of a certificate's lifetime for which a cached
// copy is reused. This pins the comparison, which is the behaviour change of
// the fix: at 70, a certificate 60% through its life is still served and one
// 80% through is dropped; at 100 it is served until it expires.
func TestCacheReadHonoursCacheForRatio(t *testing.T) {
	tcases := []struct {
		Name    string
		Ratio   int
		Elapsed float64
		Kept    bool
	}{
		{Name: "70: 60% elapsed is kept", Ratio: 70, Elapsed: 0.60, Kept: true},
		{Name: "70: 80% elapsed is dropped", Ratio: 70, Elapsed: 0.80, Kept: false},
		{Name: "100: 99% elapsed is kept", Ratio: 100, Elapsed: 0.99, Kept: true},
		{Name: "30: 50% elapsed is dropped", Ratio: 30, Elapsed: 0.50, Kept: false},
	}

	for _, tc := range tcases {
		t.Run(tc.Name, func(t *testing.T) {
			storage := &logical.InmemStorage{}
			cache := NewCache()
			ce := &CacheEntry{Users: 1, Account: "lenstra", Cert: certificateAtLifetime(t, tc.Elapsed)}
			require.NoError(t, ce.Save(context.Background(), storage, "cache/entry"))

			got, err := cache.Read(context.Background(), storage, &role{CacheForRatio: tc.Ratio}, "cache/entry")
			require.NoError(t, err)
			if tc.Kept {
				require.NotNil(t, got, "the entry should still be served")
				require.Equal(t, 2, got.Users, "a served entry gains a user")
			} else {
				require.Nil(t, got, "the entry should have been dropped")
				stored, err := storage.Get(context.Background(), "cache/entry")
				require.NoError(t, err)
				require.Nil(t, stored, "a dropped entry is removed from storage")
			}
		})
	}
}
