// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/go-acme/lego/v5/certcrypto"
	"github.com/go-acme/lego/v5/certificate"
	"github.com/go-acme/lego/v5/challenge/dns01"
	"github.com/go-acme/lego/v5/lego"
	"github.com/go-acme/lego/v5/providers/dns"
	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/logical"
)

// dnsChallengeMu serialises DNS-01 issuance. lego v5 reads the resolvers
// for its propagation checks from a process-wide default client, which
// setupChallengeProviders sets per account; two accounts with different
// resolvers issuing at once would otherwise check against each other's.
var dnsChallengeMu sync.Mutex

func getCertFromACMEProvider(ctx context.Context, logger log.Logger, req *logical.Request, a *account, names []string, keyType certcrypto.KeyType) (*certificate.Resource, error) {
	if a.Provider != "" {
		dnsChallengeMu.Lock()
		defer dnsChallengeMu.Unlock()
	}

	client, err := a.getClient()
	if err != nil {
		return nil, err
	}

	err = setupChallengeProviders(ctx, logger, client, a, req)
	if err != nil {
		return nil, err
	}

	request := certificate.ObtainRequest{
		Domains: names,
		Bundle:  true,
		// lego v5 has no default; the role decides, and a role that has not
		// decided gets the RSA 2048 that v4 issued unless told otherwise.
		KeyType: keyType,
		// v4 put the first requested name in the Subject CN; v5 leaves the CN
		// empty unless asked. Keep it, for whatever reads the CN.
		EnableCommonName: true,
	}

	return client.Certificate.Obtain(ctx, request)
}

func setupChallengeProviders(ctx context.Context, logger log.Logger, client *lego.Client, a *account, req *logical.Request) error {
	// DNS-01
	if a.Provider != "" {
		// Set the provider configuration as environment variables, so that they get picked up when the challenge provider
		// is created
		for key, value := range a.ProviderConfiguration {
			err := os.Setenv(key, value)
			if err != nil {
				return fmt.Errorf(`error setting key "%s" for DNS provider configuation`, key)
			}
		}

		provider, err := dns.NewDNSChallengeProviderByName(a.Provider)
		if err != nil {
			return err
		}

		// lego v5 takes the resolvers for propagation checks from a
		// process-wide default client rather than from a per-provider
		// option, so set it for this account, and back to the system
		// resolvers for an account that names none.
		if len(a.DNSResolvers) > 0 {
			dns01.SetDefaultClient(dns01.NewClient(&dns01.Options{RecursiveNameservers: a.DNSResolvers}))
		} else {
			dns01.SetDefaultClient(dns01.NewClient(nil))
		}

		err = client.Challenge.SetDNS01Provider(
			provider,
			dns01.CondOptions(a.IgnoreDNSPropagation, dns01.DisableAuthoritativeNssPropagationRequirement()),
		)
		if err != nil {
			return err
		}
	}

	// HTTP-01
	if a.EnableHTTP01 {
		provider := newVaultHTTP01Provider(ctx, logger, req)
		err := client.Challenge.SetHTTP01Provider(provider)
		if err != nil {
			return err
		}
	}

	// TLS-ALPN-01
	if a.EnableTLSALPN01 {
		provider := newVaultTLSALPN01Provider(ctx, logger, req)
		err := client.Challenge.SetTLSALPN01Provider(provider)
		if err != nil {
			return err
		}
	}

	return nil
}
