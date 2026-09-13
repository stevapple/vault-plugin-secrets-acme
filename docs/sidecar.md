# ACME Sidecar

While the Vault ACME secrets backend can natively solve the DNS-01 challenge when
requesting certificates, a sidecar is needed to solve both the HTTP-01 and the
TLS-ALPN-01 challenges.

## Operation

![ACME Sidecar](img/acme-sidecar.svg)
