package acme

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"github.com/go-acme/lego/v5/certificate"

	keystore "github.com/pavlo-v-chernykh/keystore-go/v4"
	"software.sslmate.com/src/go-pkcs12"
)

// Rendering formats, named as Vault's PKI engine names them so a consumer can
// ask both engines for the same thing.
const (
	formatPEM          = "pem"
	formatPKCS12Bundle = "pkcs12_bundle"
	formatJKSBundle    = "jks_bundle"

	pkcs12EncoderModern2023 = "modern2023"
	pkcs12EncoderModern2026 = "modern2026"

	defaultKeystorePassword = "changeit"
	defaultJKSAlias         = "1"
)

func allowedFormats() []interface{} {
	return []interface{}{formatPEM, formatPKCS12Bundle, formatJKSBundle}
}

func allowedPKCS12Encoders() []interface{} {
	return []interface{}{pkcs12EncoderModern2026, pkcs12EncoderModern2023}
}

// bundleOptions is how a request asks for the certificate to be rendered. None
// of it changes which certificate gets issued.
type bundleOptions struct {
	Format         string
	PKCS12Password string
	PKCS12Encoder  string
	JKSPassword    string
	JKSAlias       string
}

// validate rejects a format or encoder the engine cannot render. It runs
// before anything is looked up or ordered, so a typo in the request costs
// neither an ACME order nor a cache slot and comes back as a client error.
func (o bundleOptions) validate() error {
	switch o.Format {
	case "", formatPEM, formatPKCS12Bundle, formatJKSBundle:
	default:
		return fmt.Errorf("%q is not a supported format", o.Format)
	}
	switch o.PKCS12Encoder {
	case "", pkcs12EncoderModern2026, pkcs12EncoderModern2023:
	default:
		return fmt.Errorf("%q is not a supported pkcs12_encoder", o.PKCS12Encoder)
	}

	return nil
}

// parsedCert is a lego certificate resource taken apart into the pieces every
// bundle format needs.
type parsedCert struct {
	Leaf       *x509.Certificate
	CAChain    []*x509.Certificate
	PrivateKey interface{}
	KeyDER     []byte
}

func parseCertResource(certPEM, keyPEM []byte) (*parsedCert, error) {
	certs, err := parsePEMCertificates(certPEM)
	if err != nil {
		return nil, err
	}
	if len(certs) == 0 {
		return nil, errors.New("no certificate found in the issued PEM data")
	}

	block, _ := pem.Decode(keyPEM)
	if block == nil {
		return nil, errors.New("no PEM block found in the private key")
	}
	key, err := parsePrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// JKS stores the key as PKCS#8 DER whatever it was issued as.
	keyDER, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the private key as PKCS#8: %w", err)
	}

	return &parsedCert{Leaf: certs[0], CAChain: certs[1:], PrivateKey: key, KeyDER: keyDER}, nil
}

func parsePEMCertificates(data []byte) ([]*x509.Certificate, error) {
	var certs []*x509.Certificate
	for {
		var block *pem.Block
		block, data = pem.Decode(data)
		if block == nil {
			return certs, nil
		}
		if block.Type != "CERTIFICATE" {
			continue
		}
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse the certificate: %w", err)
		}
		certs = append(certs, cert)
	}
}

// lego hands back whatever the ACME account's key type produced, so all three
// encodings have to be tried.
func parsePrivateKey(der []byte) (interface{}, error) {
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err := x509.ParseECPrivateKey(der); err == nil {
		return key, nil
	}

	return nil, errors.New("failed to parse the private key as PKCS#8, PKCS#1 or SEC 1")
}

func encodePKCS12(c *parsedCert, opts bundleOptions) ([]byte, error) {
	var encoder *pkcs12.Encoder
	switch opts.PKCS12Encoder {
	case pkcs12EncoderModern2026, "":
		encoder = pkcs12.Modern2026
	case pkcs12EncoderModern2023:
		encoder = pkcs12.Modern2023
	default:
		return nil, fmt.Errorf("%q is not a supported pkcs12_encoder", opts.PKCS12Encoder)
	}

	return encoder.Encode(c.PrivateKey, c.Leaf, c.CAChain, opts.PKCS12Password)
}

func encodeJKS(c *parsedCert, opts bundleOptions) ([]byte, error) {
	chain := make([]keystore.Certificate, 0, len(c.CAChain)+1)
	for _, cert := range append([]*x509.Certificate{c.Leaf}, c.CAChain...) {
		chain = append(chain, keystore.Certificate{Type: "X.509", Content: cert.Raw})
	}

	// Aliases are case sensitive, and keystore-go folds them to lower case
	// unless told otherwise.
	ks := keystore.New(keystore.WithCaseExactAliases())
	entry := keystore.PrivateKeyEntry{
		CreationTime:     time.Now().UTC(),
		PrivateKey:       c.KeyDER,
		CertificateChain: chain,
	}
	password := []byte(opts.JKSPassword)
	if err := ks.SetPrivateKeyEntry(opts.JKSAlias, entry, password); err != nil {
		return nil, fmt.Errorf("failed to add the private key entry: %w", err)
	}

	var buf bytes.Buffer
	if err := ks.Store(&buf, password); err != nil {
		return nil, fmt.Errorf("failed to write the keystore: %w", err)
	}

	return buf.Bytes(), nil
}

// renderCert builds the format-dependent part of the response. A keystore
// carries the key inside the archive, so private_key and issuer_cert have
// nothing left to say and are omitted rather than returned empty.
func renderCert(cert *certificate.Resource, opts bundleOptions) (map[string]interface{}, error) {
	if opts.Format == "" || opts.Format == formatPEM {
		return map[string]interface{}{
			"private_key": string(cert.PrivateKey),
			"cert":        string(cert.Certificate),
			"issuer_cert": string(cert.IssuerCertificate),
		}, nil
	}

	if len(cert.PrivateKey) == 0 {
		return nil, fmt.Errorf("cannot build a %s without a private key", opts.Format)
	}
	parsed, err := parseCertResource(cert.Certificate, cert.PrivateKey)
	if err != nil {
		return nil, err
	}

	var archive []byte
	switch opts.Format {
	case formatPKCS12Bundle:
		archive, err = encodePKCS12(parsed, opts)
	case formatJKSBundle:
		archive, err = encodeJKS(parsed, opts)
	default:
		return nil, fmt.Errorf("%q is not a supported format", opts.Format)
	}
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"cert": base64.StdEncoding.EncodeToString(archive),
	}, nil
}
