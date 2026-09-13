// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

func TestListRoles(t *testing.T) {
	config, b := getTestBackend(t)

	listReq := &logical.Request{
		Operation: logical.ListOperation,
		Path:      "roles",
		Storage:   config.StorageView,
	}
	listResp := makeRequest(t, b, listReq, "")
	require.Equal(t, map[string]interface{}{}, listResp.Data)

	makeRequest(t, b, &logical.Request{
		Operation: logical.CreateOperation,
		Path:      "roles/lenstra",
		Storage:   config.StorageView,
		Data: map[string]interface{}{
			"account": "lenstra",
		},
	}, "")

	listResp = makeRequest(t, b, listReq, "")
	require.Equal(t, map[string]interface{}{
		"keys": []string{"lenstra"},
	}, listResp.Data)
}

// Only 100 is flagged: it is the one value whose meaning the engine can judge
// on its own, since it serves a certificate for the whole of its lifetime.
// Anything below depends on what the consumers do, which the engine cannot
// know, so it must stay accepted and unremarked.
func TestCacheForRatioServesUntilExpiryWarning(t *testing.T) {
	tcases := []struct {
		Ratio   int
		Warning bool
	}{
		{Ratio: 70, Warning: false},
		{Ratio: 90, Warning: false},
		{Ratio: 99, Warning: false},
		{Ratio: 100, Warning: true},
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("ratio %d", tc.Ratio), func(t *testing.T) {
			config := logical.TestBackendConfig()
			config.StorageView = &logical.InmemStorage{}
			b, err := Factory("test")(context.Background(), config)
			require.NoError(t, err)

			resp, err := b.HandleRequest(context.Background(), &logical.Request{
				Operation: logical.CreateOperation,
				Path:      "roles/lenstra.fr",
				Storage:   config.StorageView,
				Data: map[string]interface{}{
					"account":         "lenstra",
					"cache_for_ratio": tc.Ratio,
				},
			})
			require.NoError(t, err)
			require.NotNil(t, resp)
			require.False(t, resp.IsError())
			require.Equal(t, tc.Ratio, resp.Data["cache_for_ratio"])

			if tc.Warning {
				require.Len(t, resp.Warnings, 1)
				require.Contains(t, resp.Warnings[0], "cache_for_ratio")
			} else {
				require.Empty(t, resp.Warnings)
			}
		})
	}
}
