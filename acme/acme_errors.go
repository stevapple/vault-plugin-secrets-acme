package acme

import (
	"errors"

	"github.com/go-acme/lego/v5/acme"
)

// acmeAlreadyRevoked is the RFC 8555 problem type a server returns when asked
// to revoke a certificate that is already revoked.
const acmeAlreadyRevoked = "urn:ietf:params:acme:error:alreadyRevoked"

// alreadyRevoked reports whether a revocation failed only because the
// certificate was already revoked. For the engine that is the outcome it
// wanted: a certificate withdrawn through acme/revoke, by another ACME client
// on the same account, or by the CA itself is not a certificate to keep
// retrying against, and a lease whose revocation "fails" this way would
// otherwise sit in Vault's retry queue for good.
func alreadyRevoked(err error) bool {
	var problem *acme.ProblemDetails
	return errors.As(err, &problem) && problem.Type == acmeAlreadyRevoked
}
