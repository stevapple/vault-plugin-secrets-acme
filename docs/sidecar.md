# ACME Sidecar

While the Vault ACME secrets backend can natively solve the DNS-01 challenge when
requesting certificates, a sidecar is needed to solve both the HTTP-01 and the
TLS-ALPN-01 challenges.

This covers certificates for IP addresses as well as for domain names: solving
TLS-ALPN-01 for an address, the ACME provider sends the reverse-DNS form of it
(`1.0.0.10.in-addr.arpa`) as the TLS server name, and the sidecar reads it back
as the address to find the challenge the engine stored.

## Operation

![ACME Sidecar](img/acme-sidecar.svg)
