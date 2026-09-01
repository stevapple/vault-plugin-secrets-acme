# Security policy

This is a Vault secrets engine. It holds ACME account keys, the credentials the
DNS provider is driven with, and the private keys of the certificates it
issues. Treat anything that
discloses those, or that lets a role obtain a certificate outside its
`allowed_domains`, as a security issue rather than a bug.

## Reporting a vulnerability

Report privately through GitHub's
[private vulnerability reporting](https://github.com/stevapple/vault-plugin-secrets-acme/security/advisories/new)
— the "Report a vulnerability" button on the Security tab. Please do not open a
public issue for something exploitable.

Include the plugin version, the Vault version, and enough of the role and
account configuration to reproduce it. Redact credentials; the parameter names
and shapes are what matter, not their values.

## Scope

This engine is a fork. A vulnerability that also affects
[Boostport/vault-plugin-secrets-acme][boostport] or
[remilapeyre/vault-acme][remilapeyre] should be reported to them as well, since
a fix here does not reach their users.

Vulnerabilities in Vault itself belong to
[HashiCorp](https://github.com/hashicorp/vault/security/policy), and ones in
[lego](https://github.com/go-acme/lego) — which performs every ACME exchange and
every DNS provider call — belong to that project.

## Supported versions

Fixes land on `main` and ship in the next release. There is no long-term
support branch, and older releases are not patched.

[boostport]: https://github.com/Boostport/vault-plugin-secrets-acme
[remilapeyre]: https://github.com/remilapeyre/vault-acme
