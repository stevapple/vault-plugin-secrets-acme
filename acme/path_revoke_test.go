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

	"github.com/go-acme/lego/v5/certificate"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

func selfSigned(t *testing.T, cn string) string {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject:      pkix.Name{CommonName: cn},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	require.NoError(t, err)

	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der}))
}

// Revoking a certificate has to take it out of the cache as well, or the
// engine carries on serving one it has just withdrawn.
func TestPurgeCachedCert(t *testing.T) {
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	raw, err := Factory("test")(context.Background(), config)
	require.NoError(t, err)
	b, ok := raw.(backend)
	require.True(t, ok)

	revoked := selfSigned(t, "revoked.lenstra.fr")
	kept := selfSigned(t, "kept.lenstra.fr")

	r := &role{Account: "lenstra"}
	for key, cert := range map[string]string{"cache/aaa": revoked, "cache/bbb": kept} {
		require.NoError(t, b.cache.Create(context.Background(), config.StorageView, r, key,
			&certificate.Resource{Certificate: []byte(cert)}))
	}

	purged, err := b.purgeCachedCert(context.Background(), config.StorageView, revoked)
	require.NoError(t, err)
	require.Equal(t, 1, purged)

	gone, err := b.cache.Read(context.Background(), config.StorageView, nil, "cache/aaa")
	require.NoError(t, err)
	require.Nil(t, gone)

	survivor, err := b.cache.Read(context.Background(), config.StorageView, nil, "cache/bbb")
	require.NoError(t, err)
	require.NotNil(t, survivor)
}

// The cache holds the whole bundle while a caller is most likely to hand back
// just the leaf, so matching cannot rely on the two strings being equal.
func TestPurgeCachedCertMatchesLeafOfABundle(t *testing.T) {
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	raw, err := Factory("test")(context.Background(), config)
	require.NoError(t, err)
	b, ok := raw.(backend)
	require.True(t, ok)

	leaf := selfSigned(t, "leaf.lenstra.fr")
	issuer := selfSigned(t, "issuer.lenstra.fr")

	require.NoError(t, b.cache.Create(context.Background(), config.StorageView,
		&role{Account: "lenstra"}, "cache/aaa",
		&certificate.Resource{Certificate: []byte(leaf + issuer)}))

	purged, err := b.purgeCachedCert(context.Background(), config.StorageView, leaf)
	require.NoError(t, err)
	require.Equal(t, 1, purged)
}

// An absent reason sends none; a valid code is passed through; 7 and anything
// above 10 are not RFC 5280 reason codes and are refused as client errors.
func TestRevocationReason(t *testing.T) {
	none, err := revocationReason(-1)
	require.NoError(t, err)
	require.Nil(t, none)

	for _, code := range []int{0, 1, 3, 4, 5, 6, 8, 9, 10} {
		r, err := revocationReason(code)
		require.NoError(t, err)
		require.Equal(t, uint(code), *r)
	}
	for _, code := range []int{7, 11, 999} {
		_, err := revocationReason(code)
		require.ErrorContains(t, err, "not an RFC 5280 revocation reason")
	}
}
