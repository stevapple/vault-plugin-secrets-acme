// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
	"github.com/stretchr/testify/require"
)

func TestListRoles(t *testing.T) {
	config, b := getTestConfig(t)

	listReq := &logical.Request{
		Operation: logical.ListOperation,
		Path:      "roles",
		Storage:   config.StorageView,
	}
	listResp := makeRequest(t, b, listReq, "")
	require.Equal(t, map[string]interface{}{}, listResp.Data)

	makeRequest(t, b, &logical.Request{
		Operation: logical.CreateOperation,
		Path:      "accounts/lenstra",
		Storage:   config.StorageView,
		Data: map[string]interface{}{
			"server_url":              "https://localhost:14000/dir",
			"contact":                 "remi@lenstra.fr",
			"terms_of_service_agreed": true,
			"provider":                "exec",
		},
	}, "")
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
