package acme

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/go-acme/lego/v5/certificate"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	keystore "github.com/pavlo-v-chernykh/keystore-go/v4"
	"github.com/stretchr/testify/require"
	"software.sslmate.com/src/go-pkcs12"
)

func issuedResource(t *testing.T) *certificate.Resource {
	t.Helper()

	caKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	caTmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "Test CA"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, &caKey.PublicKey, caKey)
	require.NoError(t, err)
	ca, err := x509.ParseCertificate(caDER)
	require.NoError(t, err)

	leafKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	leafTmpl := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "www.lenstra.fr"},
		DNSNames:     []string{"www.lenstra.fr"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(24 * time.Hour),
	}
	leafDER, err := x509.CreateCertificate(rand.Reader, leafTmpl, ca, &leafKey.PublicKey, caKey)
	require.NoError(t, err)

	keyDER, err := x509.MarshalPKCS8PrivateKey(leafKey)
	require.NoError(t, err)

	pemOf := func(t string, b []byte) []byte { return pem.EncodeToMemory(&pem.Block{Type: t, Bytes: b}) }

	return &certificate.Resource{
		Domains:           []string{"www.lenstra.fr"},
		Certificate:       append(pemOf("CERTIFICATE", leafDER), pemOf("CERTIFICATE", caDER)...),
		IssuerCertificate: pemOf("CERTIFICATE", caDER),
		PrivateKey:        pemOf("PRIVATE KEY", keyDER),
	}
}

func TestRenderCertPEM(t *testing.T) {
	res := issuedResource(t)

	data, err := renderCert(res, bundleOptions{Format: formatPEM})
	require.NoError(t, err)
	require.Equal(t, string(res.Certificate), data["cert"])
	require.Equal(t, string(res.PrivateKey), data["private_key"])
	require.Equal(t, string(res.IssuerCertificate), data["issuer_cert"])
}

func TestRenderCertPKCS12(t *testing.T) {
	res := issuedResource(t)

	for _, encoder := range []string{pkcs12EncoderModern2026, pkcs12EncoderModern2023} {
		t.Run(encoder, func(t *testing.T) {
			data, err := renderCert(res, bundleOptions{
				Format:         formatPKCS12Bundle,
				PKCS12Encoder:  encoder,
				PKCS12Password: "s3cret",
			})
			require.NoError(t, err)

			// The archive carries the key, so these have nothing left to say.
			require.NotContains(t, data, "private_key")
			require.NotContains(t, data, "issuer_cert")

			archive, err := base64.StdEncoding.DecodeString(data["cert"].(string))
			require.NoError(t, err)

			key, leaf, chain, err := pkcs12.DecodeChain(archive, "s3cret")
			require.NoError(t, err)
			require.NotNil(t, key)
			require.Equal(t, "www.lenstra.fr", leaf.Subject.CommonName)
			require.Len(t, chain, 1)
			require.Equal(t, "Test CA", chain[0].Subject.CommonName)
		})
	}
}

func TestRenderCertJKS(t *testing.T) {
	res := issuedResource(t)

	data, err := renderCert(res, bundleOptions{
		Format:      formatJKSBundle,
		JKSPassword: "s3cret",
		JKSAlias:    "Server",
	})
	require.NoError(t, err)
	require.NotContains(t, data, "private_key")

	archive, err := base64.StdEncoding.DecodeString(data["cert"].(string))
	require.NoError(t, err)

	ks := keystore.New(keystore.WithCaseExactAliases())
	require.NoError(t, ks.Load(bytes.NewReader(archive), []byte("s3cret")))
	// Case preserved, as Vault documents the alias to be case sensitive.
	require.Equal(t, []string{"Server"}, ks.Aliases())

	entry, err := ks.GetPrivateKeyEntry("Server", []byte("s3cret"))
	require.NoError(t, err)
	require.Len(t, entry.CertificateChain, 2)
	_, err = x509.ParsePKCS8PrivateKey(entry.PrivateKey)
	require.NoError(t, err)
}

func TestRenderCertRejectsUnknownFormat(t *testing.T) {
	_, err := renderCert(issuedResource(t), bundleOptions{Format: "pkcs7"})
	require.ErrorContains(t, err, "not a supported format")
}

// Rendering options are not part of what gets ordered from the ACME provider,
// so two requests that differ only in format must share one cache entry rather
// than each issuing a certificate of their own.
func TestCacheKeyIgnoresRenderingOptions(t *testing.T) {
	r := &role{Account: "lenstra", AllowedDomains: []string{"lenstra.fr"}, CacheForRatio: 70}
	schema := pathCerts(&backend{}).Fields

	keyFor := func(raw map[string]interface{}) string {
		t.Helper()
		k, err := getCacheKey(r, &framework.FieldData{Raw: raw, Schema: schema})
		require.NoError(t, err)
		return k
	}

	base := keyFor(map[string]interface{}{"role": "lenstra.fr", "common_name": "www.lenstra.fr"})

	for _, raw := range []map[string]interface{}{
		{"role": "lenstra.fr", "common_name": "www.lenstra.fr", "format": formatPKCS12Bundle},
		{"role": "lenstra.fr", "common_name": "www.lenstra.fr", "format": formatJKSBundle, "jks_password": "other"},
		{"role": "lenstra.fr", "common_name": "www.lenstra.fr", "pkcs12_encoder": pkcs12EncoderModern2023},
		{"role": "lenstra.fr", "common_name": "www.lenstra.fr", "jks_private_key_alias": "Server"},
	} {
		require.Equal(t, base, keyFor(raw), "rendering options must not split the cache")
	}

	// A different name still must.
	require.NotEqual(t, base, keyFor(map[string]interface{}{"role": "lenstra.fr", "common_name": "other.lenstra.fr"}))
}

// An unsupported format is a client error, and it must be reported before the
// request gets anywhere near the cache or the ACME provider: here there is not
// even a role to find, and the format is still what the response complains
// about.
func TestCertCreateRejectsUnknownFormatFirst(t *testing.T) {
	config := logical.TestBackendConfig()
	config.StorageView = &logical.InmemStorage{}
	b, err := Factory("test")(context.Background(), config)
	require.NoError(t, err)

	for _, data := range []map[string]interface{}{
		{"common_name": "www.lenstra.fr", "format": "pkcs7"},
		{"common_name": "www.lenstra.fr", "format": formatPKCS12Bundle, "pkcs12_encoder": "legacy"},
	} {
		resp, err := b.HandleRequest(context.Background(), &logical.Request{
			Operation: logical.CreateOperation,
			Path:      "certs/lenstra.fr",
			Storage:   config.StorageView,
			Data:      data,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.True(t, resp.IsError())
		require.Contains(t, resp.Error().Error(), "is not a supported")
	}
}
