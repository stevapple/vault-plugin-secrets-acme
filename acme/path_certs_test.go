// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

func TestCerts(t *testing.T) {
	config, b := getTestConfig(t)
	createAccount(t, b, config.StorageView)
	createRole(t, b, config.StorageView)

	t.Log("Try to request certificates")
	firstCert, secondCert := checkCreatingCerts(t, b, config.StorageView)

	// Try to run an HTTPS server with the certificate we got and query it
	checkCertificate(t, firstCert)

	t.Log("Try to renew the lease")
	checkRenewingCert(t, b, config.StorageView, firstCert.Secret)

	t.Log("Try to revoke the lease")
	checkRevokeCert(t, b, config.StorageView, firstCert, secondCert)
}

func TestExplicitProviderConfiguration(t *testing.T) {
	config, b := getTestConfig(t)
	req := &logical.Request{
		Operation: logical.CreateOperation,
		Path:      "accounts/lenstra",
		Storage:   config.StorageView,
		Data: map[string]interface{}{
			"server_url":              "https://localhost:14000/dir",
			"contact":                 "remi@lenstra.fr",
			"terms_of_service_agreed": true,
			"provider":                "exec",
			"dns_resolvers":           []string{"127.0.0.1:8053"},
			"provider_configuration": map[string]string{
				// We use a bad configuration on purpose so we get a failure
				// when requesting a certificate instead of the success we got
				// in the previous test
				"EXEC_PATH": "/dev/null",
			},
		},
	}
	makeRequest(t, b, req, "")
	createRole(t, b, config.StorageView)

	req = &logical.Request{
		Operation: logical.CreateOperation,
		Path:      "certs/lenstra.fr",
		Storage:   config.StorageView,
		Data:      map[string]interface{}{"common_name": "sentry.lenstra.fr"},
	}
	resp, err := b.HandleRequest(context.Background(), req)
	require.Error(t, err, "fork/exec /dev/null: permission denied")
	require.Equal(t, resp.Data, map[string]interface{}{
		"error": "Failed to validate certificate signing request: resolver: one or more domains had a problem: [sentry.lenstra.fr: dns01: error presenting token (sentry.lenstra.fr): exec: start command: fork/exec /dev/null: permission denied]",
	})
}

func checkCreatingCerts(t *testing.T, b logical.Backend, storage logical.Storage) (*logical.Response, *logical.Response) {
	certReq := &logical.Request{
		Operation: logical.CreateOperation,
		Path:      "certs/foo",
		Storage:   storage,
		Data:      map[string]interface{}{"common_name": "sentry.lenstra.fr"},
	}
	makeRequest(t, b, certReq, "This role does not exist")

	// Try with an existing role
	certReq.Path = "certs/lenstra.fr"
	makeRequest(t, b, certReq, "")

	// Try with alternate names
	certReq.Data = map[string]interface{}{
		"common_name":       "sentry.lenstra.fr",
		"alternative_names": "grafana.lenstra.fr",
	}
	first := makeRequest(t, b, certReq, "")
	second := makeRequest(t, b, certReq, "")

	// Since caching is enabled, we should get the same cert when calling the
	// endpoint twice
	require.Equal(t, first.Data, second.Data)

	// The lease spans the certificate it carries.
	assertLeaseSpansCertificate(t, first)
	assertLeaseSpansCertificate(t, second)

	return first, second
}

// assertLeaseSpansCertificate checks that both the lease's TTL and its MaxTTL
// run to the certificate's NotAfter, within a minute for the time the request
// took.
func assertLeaseSpansCertificate(t *testing.T, resp *logical.Response) {
	t.Helper()

	block, _ := pem.Decode([]byte(resp.Data["cert"].(string)))
	require.NotNil(t, block)
	leaf, err := x509.ParseCertificate(block.Bytes)
	require.NoError(t, err)
	require.NotNil(t, resp.Secret)

	validity := time.Until(leaf.NotAfter).Seconds()
	require.InDelta(t, validity, resp.Secret.TTL.Seconds(), 60)
	require.InDelta(t, validity, resp.Secret.MaxTTL.Seconds(), 60)
}

func checkRenewingCert(t *testing.T, b logical.Backend, storage logical.Storage, secret *logical.Secret) {
	certReq := &logical.Request{
		Operation: logical.RenewOperation,
		Path:      "certs/lenstra.fr",
		Storage:   storage,
		Data:      map[string]interface{}{"common_name": "sentry.lenstra.fr"},
		Secret:    secret,
	}
	renewResp := makeRequest(t, b, certReq, "")

	if renewResp.Secret.TTL < secret.TTL {
		t.Fatalf("Failed to renew secret")
	}
}

func checkRevokeCert(t *testing.T, b logical.Backend, storage logical.Storage, first, second *logical.Response) {
	certReq := &logical.Request{
		Operation: logical.RevokeOperation,
		Path:      "certs/lenstra.fr",
		Storage:   storage,
		Secret:    first.Secret,
	}
	makeRequest(t, b, certReq, "")

	certReq.Secret = second.Secret
	makeRequest(t, b, certReq, "")

	// Check the cert status
	a, err := getAccount(context.Background(), storage, "accounts/lenstra")
	if err != nil {
		t.Fatal(err)
	}
	if a == nil {
		t.Fatal("Account should have been found")
	}
	client, err := a.getClient()
	if err != nil {
		t.Fatal(err)
	}

	// Checking the OCSP status was not working for tests
	err = client.Certificate.Revoke(context.Background(), []byte(second.Data["cert"].(string)))
	if err == nil {
		t.Fatalf("Trying to revoke the cert should have failed")
	}
	if !strings.Contains(err.Error(), "Certificate has already been revoked.") {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func checkCertificate(t *testing.T, resp *logical.Response) {
	mux := http.NewServeMux()
	mux.Handle(
		"/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte("Hello world\n"))
		}),
	)
	server := &http.Server{Addr: ":4443", Handler: mux}

	config := &tls.Config{}
	config.NextProtos = []string{"http/1.1"}

	cert, err := tls.X509KeyPair(
		[]byte(resp.Data["cert"].(string)),
		[]byte(resp.Data["private_key"].(string)),
	)
	if err != nil {
		t.Fatal(err)
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0] = cert

	ln, err := net.Listen("tcp", "0.0.0.0:4443")
	if err != nil {
		t.Fatal(err)
	}

	tlsListener := tls.NewListener(ln.(*net.TCPListener), config)

	go func() {
		err := server.Serve(tlsListener)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Error(err)
		}
	}()
	t.Cleanup(func() { _ = server.Close() })

	// Verify with Go's own verifier and a root pool of our own, rather than
	// whatever the platform verifier makes of an untrusted chain: the errors
	// it produces are typed, and the same on every OS. Pebble generates its
	// issuing root at startup and publishes it on the management interface,
	// which is itself served under the fixed test certificate in test/certs.
	roots := x509.NewCertPool()
	minica, err := os.ReadFile("../test/certs/pebble.minica.pem")
	require.NoError(t, err)
	require.True(t, roots.AppendCertsFromPEM(minica))

	management := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{RootCAs: roots}}}
	rootResp, err := management.Get("https://localhost:15000/roots/0")
	require.NoError(t, err)
	rootPEM, err := io.ReadAll(rootResp.Body)
	_ = rootResp.Body.Close()
	require.NoError(t, err)
	require.True(t, roots.AppendCertsFromPEM(rootPEM), "Pebble's root should be PEM")

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: roots},
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if addr == "example.com:443" || addr == "sentry.lenstra.fr:443" {
				addr = "127.0.0.1:4443"
			}
			return dialer.DialContext(ctx, network, addr)
		},
	}}

	// The certificate is for the names that were requested, and nothing else.
	_, err = client.Get("https://example.com")
	var hostnameErr x509.HostnameError
	require.ErrorAs(t, err, &hostnameErr)
	require.Equal(t, "example.com", hostnameErr.Host)

	// And it chains to the CA that issued it.
	httpResp, err := client.Get("https://sentry.lenstra.fr")
	require.NoError(t, err)
	body, err := io.ReadAll(httpResp.Body)
	_ = httpResp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, "Hello world\n", string(body))
}
