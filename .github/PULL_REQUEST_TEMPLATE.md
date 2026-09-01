## What this changes

<!-- And why. The diff says what; the reason it is being made is the part a
reader cannot reconstruct. -->

## Effect on existing deployments

<!-- Delete if the change cannot reach one. Otherwise, answer both:

  - What does a consumer already holding a certificate see? Certificates
    outlive the leases that fetched them, so issuance, caching, lease and
    revocation changes reach places the request never did.
  - Does an existing role or mount behave differently after upgrading with no
    configuration change? If so, say so plainly — it decides whether this is a
    patch or a minor release. -->

## Checklist

- [ ] `make test` passes, or the reason it cannot run here is stated
- [ ] `golangci-lint run` is clean
- [ ] Documentation updated, or not applicable
- [ ] If a parameter was added, it is spelled and defaulted as Vault's PKI
      engine spells and defaults the equivalent
