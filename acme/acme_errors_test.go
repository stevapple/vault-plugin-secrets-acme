package acme

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-acme/lego/v5/acme"
	"github.com/stretchr/testify/require"
)

func TestAlreadyRevoked(t *testing.T) {
	already := &acme.ProblemDetails{Type: acmeAlreadyRevoked, HTTPStatus: 400}
	other := &acme.ProblemDetails{Type: "urn:ietf:params:acme:error:unauthorized", HTTPStatus: 403}

	require.True(t, alreadyRevoked(already))
	require.True(t, alreadyRevoked(fmt.Errorf("acme: %w", already)), "lego wraps the problem, so unwrapping must work")
	require.False(t, alreadyRevoked(other))
	require.False(t, alreadyRevoked(errors.New("connection refused")))
	require.False(t, alreadyRevoked(nil))
}
