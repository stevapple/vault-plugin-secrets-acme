// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/go-acme/lego/v5/certcrypto"
	"github.com/go-acme/lego/v5/certificate"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

func pathCerts(b *backend) *framework.Path {
	return &framework.Path{
		Pattern: "certs/" + framework.GenericNameRegex("role"),
		Fields: map[string]*framework.FieldSchema{
			"role": {
				Type:     framework.TypeString,
				Required: true,
			},
			"common_name": {
				Type:     framework.TypeString,
				Required: true,
			},
			"alternative_names": {
				Type: framework.TypeCommaStringSlice,
			},
			"format": {
				Type:          framework.TypeString,
				Default:       formatPEM,
				AllowedValues: allowedFormats(),
			},
			"pkcs12_password": {
				Type:    framework.TypeString,
				Default: defaultKeystorePassword,
			},
			"pkcs12_encoder": {
				Type:          framework.TypeString,
				Default:       pkcs12EncoderModern2026,
				AllowedValues: allowedPKCS12Encoders(),
			},
			"jks_password": {
				Type:    framework.TypeString,
				Default: defaultKeystorePassword,
			},
			"jks_private_key_alias": {
				Type:    framework.TypeString,
				Default: defaultJKSAlias,
			},
		},
		ExistenceCheck: b.pathExistenceCheck,
		Operations: map[logical.Operation]framework.OperationHandler{
			logical.CreateOperation: &framework.PathOperation{
				Callback: b.certCreate,
			},
		},
	}
}

func (b *backend) certCreate(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	opts := bundleOptions{
		Format:         data.Get("format").(string),
		PKCS12Password: data.Get("pkcs12_password").(string),
		PKCS12Encoder:  data.Get("pkcs12_encoder").(string),
		JKSPassword:    data.Get("jks_password").(string),
		JKSAlias:       data.Get("jks_private_key_alias").(string),
	}
	if err := opts.validate(); err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	names := getNames(data)

	rolePath := "roles/" + data.Get("role").(string)
	r, err := getRole(ctx, req.Storage, rolePath)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return logical.ErrorResponse("This role does not exist"), nil
	}
	if err = validateNames(b, r, names); err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}

	accountPath := "accounts/" + r.Account
	a, err := getAccount(ctx, req.Storage, accountPath)
	if err != nil {
		return nil, err
	}
	if a == nil {
		return logical.ErrorResponse("This account does not exist"), nil
	}
	// Lookup cache
	cacheKey, err := getCacheKey(r, data)
	if err != nil {
		return nil, fmt.Errorf("failed to get cache key: %v", err)
	}

	var cert *certificate.Resource

	// Let's first check the cache to see if a cert already exists
	if !r.DisableCache {
		b.cache.Lock()
		defer b.cache.Unlock()
		b.Logger().Debug("Look in the cache for a saved cert")
		ce, err := b.cache.Read(ctx, req.Storage, r, cacheKey)
		if err != nil {
			return nil, err
		}
		if ce == nil {
			b.Logger().Debug("Certificate not found in the cache")
		} else {
			cert = ce.Certificate()
		}
	}

	// If we did not find a cert, we have to request one
	if cert == nil {
		b.Logger().Debug("Contacting the ACME provider to get a new certificate")
		keyType, err := getKeyType(r.keyTypeName())
		if err != nil {
			return nil, fmt.Errorf("role has an unusable key_type: %w", err)
		}
		cert, err = getCertFromACMEProvider(ctx, b.Logger(), req, a, names, keyType)
		if err != nil {
			return logical.ErrorResponse("Failed to validate certificate signing request: %s", err), err
		}
		// Save the cert in the cache for the next request
		if !r.DisableCache {
			err = b.cache.Create(ctx, req.Storage, r, cacheKey, cert)
			if err != nil {
				return nil, err
			}
		}
	}

	s, err := b.getSecret(accountPath, rolePath, cacheKey, cert, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create the secret: %v", err)
	}

	return s, nil
}

func getCacheKey(r *role, data *framework.FieldData) (string, error) {
	rolePath, err := json.Marshal(r)
	if err != nil {
		return "", fmt.Errorf("failed to marshall role: %v", err)
	}

	d := make(map[string]interface{})
	for key := range data.Schema {
		// How the response is rendered does not change which certificate is
		// issued, so it must not split the cache: two requests differing only
		// in format would otherwise each order one of their own.
		if renderingFields[key] {
			continue
		}
		// Two spellings of one address are one certificate.
		switch key {
		case "common_name":
			d[key] = canonicalName(data.Get(key).(string))
		case "alternative_names":
			altNames := data.Get(key).([]string)
			canonical := make([]string, len(altNames))
			for i, n := range altNames {
				canonical[i] = canonicalName(n)
			}
			d[key] = canonical
		default:
			d[key] = data.Get(key)
		}
	}
	dataPath, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("failed to marshall data: %v", err)
	}

	key := string(rolePath) + string(dataPath)
	hashedKey := sha256.Sum256([]byte(key))

	return fmt.Sprintf("%s%x", cachePrefix, hashedKey), nil
}

var renderingFields = map[string]bool{
	"format":                true,
	"pkcs12_password":       true,
	"pkcs12_encoder":        true,
	"jks_password":          true,
	"jks_private_key_alias": true,
}

func (b *backend) getSecret(accountPath, rolePath, cacheKey string, cert *certificate.Resource, opts bundleOptions) (*logical.Response, error) {
	// Use the helper to create the secret
	b.Logger().Debug("Preparing response")
	certs, err := certcrypto.ParsePEMBundle(cert.Certificate)
	if err != nil {
		return nil, err
	}

	notBefore := certs[0].NotBefore
	notAfter := certs[0].NotAfter

	data, err := renderCert(cert, opts)
	if err != nil {
		return nil, err
	}
	data["domain"] = firstDomain(cert.Domains)
	data["domains"] = cert.Domains
	data["url"] = cert.CertStableURL
	data["not_before"] = notBefore.String()
	data["not_after"] = notAfter.String()

	s := b.Secret(secretCertType).Response(
		data,
		// this will be used when revoking the certificate, so it holds the PEM
		// regardless of how the response above was rendered
		map[string]interface{}{
			"account":   accountPath,
			"role":      rolePath,
			"cert":      string(cert.Certificate),
			"url":       cert.CertStableURL,
			"cache_key": cacheKey,
		})

	// Both, so the lease spans the certificate it carries. With only MaxTTL
	// set the lease falls back to the mount's default TTL and expires long
	// before the certificate does, which strands the certificate with no lease
	// left to renew it.
	validity := time.Until(notAfter)
	s.Secret.TTL = validity
	s.Secret.MaxTTL = validity

	return s, nil
}

func getNames(data *framework.FieldData) []string {
	altNames := data.Get("alternative_names").([]string)
	names := make([]string, len(altNames)+1)
	names[0] = canonicalName(data.Get("common_name").(string))
	for i, n := range altNames {
		names[i+1] = canonicalName(n)
	}

	return names
}

// canonicalName writes an address the one way (RFC 5952 for IPv6, dotted
// quad for IPv4 however it was spelled) and leaves a domain name alone. lego
// sends an identifier exactly as written and the challenge is stored under
// it, while the sidecar decodes the address the ACME server presents in its
// canonical form; the two only meet if the request is canonical too. It also
// keeps two spellings of one address from becoming two cache entries.
func canonicalName(name string) string {
	if ip := net.ParseIP(name); ip != nil {
		return ip.String()
	}

	return name
}

func validateNames(b logical.Backend, r *role, names []string) error {
	b.Logger().Debug("Validate names", "role", r, "names", names)

	isSubdomain := func(domain, root string) bool {
		return strings.HasSuffix(domain, "."+root)
	}

	for _, name := range names {
		// An address is an identifier of its own type (RFC 8738), not a name
		// under any of the allowed domains, so allowed_domains has nothing to
		// say about it and the role's own flag decides.
		if net.ParseIP(name) != nil {
			if !r.AllowIPSANs {
				return fmt.Errorf("'%s' is an IP address and the role does not allow IP SANs", name)
			}
			continue
		}

		var valid bool
		for _, domain := range r.AllowedDomains {
			if (domain == name && r.AllowBareDomains) ||
				(isSubdomain(name, domain) && r.AllowSubdomains) {
				valid = true
			}
		}
		if !valid {
			return fmt.Errorf("'%s' is not an allowed domain", name)
		}
	}

	return nil
}
