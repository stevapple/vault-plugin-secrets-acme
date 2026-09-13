// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package sidecar

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"slices"
	"strings"

	"github.com/go-acme/lego/v5/challenge/tlsalpn01"
	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/logical"
)

type client interface {
	Read(path string) (*api.Secret, error)
}

type MockClient struct {
	backend logical.Backend
	storage logical.Storage
}

// NewMockClient return a MockClient that can be used in tetss
func NewMockClient(b logical.Backend, storage logical.Storage) *MockClient {
	return &MockClient{
		backend: b,
		storage: storage,
	}
}

func (c MockClient) Read(path string) (*api.Secret, error) {
	req := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      path,
		Storage:   c.storage,
	}
	ctx := context.Background()
	resp, err := c.backend.HandleRequest(ctx, req)

	if resp == nil {
		return nil, err
	}
	return &api.Secret{
		Data: resp.Data,
	}, err
}

// Provider solves an ACME challenge
type Provider interface {
	Listen(addr string) error
}

type http01Provider struct {
	client client
	logger log.Logger
}

type acmeHandler struct {
	client client
}

func (a acmeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.URL.Path, "/.well-known/acme-challenge/")
	path := fmt.Sprintf("challenges/http-01/%s", token)

	s, err := a.client.Read(path)
	// Nothing useful can be done about a write to the response failing: the
	// ACME server has gone away, and it will report the challenge as failed.
	if err != nil {
		_, _ = fmt.Fprintf(w, "failed to read token: %s", err.Error())
	} else if err, ok := s.Data["error"]; ok {
		_, _ = fmt.Fprintf(w, "failed to read token: %s", err)
	} else {
		headers := w.Header()
		headers.Set("host", s.Data["key"].(string))
		headers.Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		_, _ = io.WriteString(w, s.Data["key"].(string))
	}
}

// NewHTTP01Provider returns a provider that solves the HTTP-01 challenge
// defined in https://ietf-wg-acme.github.io/acme/draft-ietf-acme-acme.html#rfc.section.8.3
func NewHTTP01Provider(client client, logger log.Logger) Provider {
	return http01Provider{
		client: client,
		logger: logger,
	}
}

func (p http01Provider) Listen(addr string) error {
	handler := acmeHandler{
		client: p.client,
	}
	http.Handle("/.well-known/acme-challenge/", handler)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}
	go func() {
		if err := http.Serve(listener, nil); err != nil {
			p.logger.Error("HTTP-01 listener stopped", "addr", addr, "error", err)
		}
	}()

	return nil
}

func (p http01Provider) Close() error {
	return nil
}

type tlsALPN01Provider struct {
	client client
	logger log.Logger
}

// NewTLSALPN01Provider returns a provider that solves the TLS-ALPN-01 challenge
// defined in https://tools.ietf.org/html/draft-ietf-acme-tls-alpn-01
func NewTLSALPN01Provider(client client, logger log.Logger) Provider {
	return tlsALPN01Provider{
		client: client,
		logger: logger,
	}
}

func (p tlsALPN01Provider) Listen(addr string) error {
	tlsConfig := &tls.Config{
		NextProtos: []string{"acme-tls/1"},
		GetCertificate: func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
			path := fmt.Sprintf("challenges/tls-alpn-01/%s", extractAddressFromReverse(hello.ServerName))

			s, err := p.client.Read(path)
			if err != nil {
				p.logger.Error("failed to read token", "err", err)
				return nil, err
			}
			if err, ok := s.Data["error"]; ok {
				p.logger.Error(err.(string))
				return nil, fmt.Errorf("failed to read token: %s", err.(string))
			}

			if len(hello.SupportedProtos) != 1 || hello.SupportedProtos[0] != "acme-tls/1" {
				return nil, fmt.Errorf("the protocol is not correct")
			}

			return tlsalpn01.ChallengeCert(s.Data["domain"].(string), s.Data["key"].(string))
		},
	}

	p.logger.Info("Listening for TLS-ALPN-01 challenge", "addr", addr)
	listener, err := tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to create listener: %v", err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			// The handshake is the whole exchange: it presents the challenge
			// certificate, and a failure only means the peer was not the
			// ACME server, or gave up. Either way the connection is done.
			_ = conn.(*tls.Conn).Handshake()
			_ = conn.Close()
		}
	}()

	return nil
}

const (
	ipv4ReverseSuffix = ".in-addr.arpa"
	ipv6ReverseSuffix = ".ip6.arpa"
)

// extractAddressFromReverse turns the reverse-DNS form of an address back
// into the address itself. Validating an IP identifier (RFC 8738) over
// TLS-ALPN-01 (RFC 8737 section 4), an ACME server sends the reverse name of
// the address as the SNI - "1.0.0.10.in-addr.arpa" for 10.0.0.1, or the
// nibble form under ".ip6.arpa" for an IPv6 address - while the engine stores
// the challenge under the address itself. Anything that is not a reverse
// name, and anything that is one but does not decode to an address, is
// returned unchanged, so a domain reaches storage as itself and a malformed
// name fails the lookup it was always going to fail.
//
// The decoding follows Boostport/vault-plugin-secrets-acme commit 9dcb44d.
func extractAddressFromReverse(name string) string {
	// The reverse name is a fully qualified one, and DNS names are compared
	// without regard to case.
	trimmed := strings.ToLower(strings.TrimSuffix(name, "."))

	var ip net.IP
	switch {
	case strings.HasSuffix(trimmed, ipv4ReverseSuffix):
		ip = addressFromReverseIPv4(strings.TrimSuffix(trimmed, ipv4ReverseSuffix))
	case strings.HasSuffix(trimmed, ipv6ReverseSuffix):
		ip = addressFromReverseIPv6(strings.TrimSuffix(trimmed, ipv6ReverseSuffix))
	default:
		return name
	}
	if ip == nil {
		return name
	}

	return ip.String()
}

// addressFromReverseIPv4 reads the four octets of an IPv4 address from the
// labels of its "in-addr.arpa" name, which hold them least significant first.
func addressFromReverseIPv4(labels string) net.IP {
	octets := strings.Split(labels, ".")
	if len(octets) != net.IPv4len {
		return nil
	}
	slices.Reverse(octets)

	return net.ParseIP(strings.Join(octets, "."))
}

// addressFromReverseIPv6 reads the sixteen bytes of an IPv6 address from the
// labels of its "ip6.arpa" name, which hold one nibble each, least
// significant first (RFC 3596 section 2.5).
func addressFromReverseIPv6(labels string) net.IP {
	nibbles := strings.Split(labels, ".")
	if len(nibbles) != 2*net.IPv6len {
		return nil
	}
	slices.Reverse(nibbles)

	groups := make([]string, 0, net.IPv6len/2)
	for i := 0; i < len(nibbles); i += 4 {
		groups = append(groups, strings.Join(nibbles[i:i+4], ""))
	}

	return net.ParseIP(strings.Join(groups, ":"))
}
