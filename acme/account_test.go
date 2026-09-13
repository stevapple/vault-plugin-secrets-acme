package acme

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

// storedAccount writes an account entry the way a given version of the engine
// would have, so that getAccount can be exercised against stores written
// before a field existed.
func storedAccount(t *testing.T, storage logical.Storage, extra map[string]interface{}) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	der, err := x509.MarshalPKCS8PrivateKey(key)
	require.NoError(t, err)

	// The keys every version has written since the first release.
	data := map[string]interface{}{
		"server_url":              "https://localhost:14000/dir",
		"registration_uri":        "https://localhost:14000/my-account/1",
		"contact":                 "rem@lenstra.fr",
		"terms_of_service_agreed": true,
		"private_key":             string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: der})),
		"provider":                "exec",
		"enable_http_01":          false,
		"enable_tls_alpn_01":      false,
	}
	for k, v := range extra {
		data[k] = v
	}
	entry, err := logical.StorageEntryJSON("accounts/lenstra", data)
	require.NoError(t, err)
	require.NoError(t, storage.Put(context.Background(), entry))
}

// key_type, provider_configuration, ignore_dns_propagation and dns_resolvers
// were each added after the first release. An account written before any of
// them existed has to load, with the field at its zero value, rather than take
// the plugin down on a type assertion.
func TestGetAccountFromOlderStore(t *testing.T) {
	tcases := []struct {
		Name      string
		Extra     map[string]interface{}
		Resolvers []string
	}{
		{Name: "first release: none of the later keys", Extra: map[string]interface{}{}},
		{Name: "dns_resolvers stored as null", Extra: map[string]interface{}{"dns_resolvers": nil}},
		{Name: "dns_resolvers stored empty", Extra: map[string]interface{}{"dns_resolvers": []string{}}},
		{
			Name:      "current layout",
			Extra:     map[string]interface{}{"key_type": "EC256", "provider_configuration": map[string]string{"EXEC_PATH": "/bin/true"}, "ignore_dns_propagation": true, "dns_resolvers": []string{"127.0.0.1:8053"}},
			Resolvers: []string{"127.0.0.1:8053"},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.Name, func(t *testing.T) {
			storage := &logical.InmemStorage{}
			storedAccount(t, storage, tc.Extra)

			a, err := getAccount(context.Background(), storage, "accounts/lenstra")
			require.NoError(t, err)
			require.NotNil(t, a)
			require.Equal(t, "rem@lenstra.fr", a.Email)
			require.Equal(t, "exec", a.Provider)
			require.NotNil(t, a.Key)
			require.NotNil(t, a.ProviderConfiguration)
			require.Equal(t, len(tc.Resolvers), len(a.DNSResolvers))
			if tc.Resolvers != nil {
				require.Equal(t, tc.Resolvers, a.DNSResolvers)
				require.Equal(t, "EC256", a.KeyType)
				require.True(t, a.IgnoreDNSPropagation)
				require.Equal(t, "/bin/true", a.ProviderConfiguration["EXEC_PATH"])
			}
		})
	}
}

// An account whose stored key is not PEM at all is an error naming the account,
// not a nil dereference.
func TestGetAccountRejectsUnparsableKey(t *testing.T) {
	storage := &logical.InmemStorage{}
	entry, err := logical.StorageEntryJSON("accounts/lenstra", map[string]interface{}{
		"server_url": "https://localhost:14000/dir", "registration_uri": "x", "contact": "rem@lenstra.fr",
		"terms_of_service_agreed": true, "private_key": "not pem", "provider": "exec",
		"enable_http_01": false, "enable_tls_alpn_01": false,
	})
	require.NoError(t, err)
	require.NoError(t, storage.Put(context.Background(), entry))

	_, err = getAccount(context.Background(), storage, "accounts/lenstra")
	require.ErrorContains(t, err, "no PEM block")
}
