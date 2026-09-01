# Contributing

This repository is a fork — see [About this fork](README.md#about-this-fork) for
what it descends from. Most of the code is the upstream authors', and a change
that is not specific to this fork is worth sending to
[Boostport/vault-plugin-secrets-acme][boostport] as well, since merging it here
does not reach anyone using theirs.

## Building

```sh
$ make build
```

The `go` directive in each module's `go.mod` is the oldest Go that can build
it. The workspace pins a newer toolchain in `go.work`, and that is what CI and
the release binaries are built with; you do not need it to build or contribute.

## Tests

```sh
$ make test      # unit tests
$ make testacc   # acceptance tests
```

Most of the suite drives a real ACME exchange, so it needs
[Pebble](https://github.com/letsencrypt/pebble) and `pebble-challtestsrv` on
`$PATH`; `make testacc` also needs a `vault` binary. Tests that only touch
storage build their backend through `getTestBackend` and run without any of
that, which is worth preferring for anything that does not have to talk to an
ACME server.

CI runs the acceptance tests against the newest release of each Vault major
line. A change that only works against one of them is a finding, not a
formality.

## Lint

```sh
$ golangci-lint run
$ (cd acme/sidecar && golangci-lint run)
```

golangci-lint stops at a module boundary, so the sidecar module is linted from
its own directory; CI does both. The configured set is golangci-lint's standard
one and the build is expected to be clean against it. Adding a linter is a change to `.golangci.yml` and a
decision about the noise it brings, not something to do in passing.

## Opening a pull request

The pull request template asks two questions besides what the change is and
why. They matter more here than in most projects, so here is what they mean.

**What does a consumer that already holds a certificate see?** Answer this if
the change touches issuance, caching, leases or revocation. A certificate is
used long after the request that fetched it, often by something that never
talks to Vault again, and several requests can share one cached certificate.
So a change can reach systems that made no request at all. For example,
changing what happens when a lease expires affects every certificate already
out there, not just the next one issued.

**Does an existing role or mount behave differently after the upgrade, with no
configuration change?** Say so plainly if it does; it decides whether the
release is a patch or a minor. For example, fixing how `cache_for_ratio` is
interpreted changes how long every existing role reuses a certificate, even
though no role's value changed.

If the change cannot reach an existing deployment at all, drop that section of
the template, as it says.

## Commit messages

Write the subject in the imperative and say what the change does. Use the body
for why it is being made, and for anything a reader would otherwise have to
work out from the diff.

[boostport]: https://github.com/Boostport/vault-plugin-secrets-acme
