// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package acme

import (
	"context"

	"github.com/go-viper/mapstructure/v2"
	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

// At 100 a cached certificate is served for the whole of its lifetime, so the
// last request before it drops out of the cache receives one with no usable
// life left. Where a consumer rotates below that is its own business and not
// something the engine can know, so this warns rather than refuses.
const cacheForRatioServesUntilExpiry = 100

func pathRoles(b *backend) []*framework.Path {
	return []*framework.Path{
		{
			Pattern: "roles/?$",
			Operations: map[logical.Operation]framework.OperationHandler{
				logical.ListOperation: &framework.PathOperation{
					Callback: b.roleList,
				},
			},
		},
		{
			Pattern: "roles/" + framework.GenericNameRegex("role"),
			Fields: map[string]*framework.FieldSchema{
				"account": {
					Type:     framework.TypeString,
					Required: true,
				},
				"allowed_domains": {
					Type: framework.TypeCommaStringSlice,
				},
				"allow_bare_domains": {
					Type: framework.TypeBool,
				},
				"allow_subdomains": {
					Type: framework.TypeBool,
				},
				"disable_cache": {
					Type: framework.TypeBool,
				},
				"cache_for_ratio": {
					Type:    framework.TypeInt,
					Default: 70,
				},
				"revoke_on_lease_expiry": {
					Type:        framework.TypeBool,
					Description: "Revoke the certificate at the ACME provider once the last lease on it goes away. Off by default: a lease expiring is not on its own evidence that the certificate has stopped being used.",
				},
				"key_type": {
					Type:          framework.TypeString,
					Default:       defaultCertKeyType,
					AllowedValues: keyTypes,
					Description:   "Type of the private key generated for each certificate, named as for an account's key_type. RSA2048 by default, which is what every certificate got before the setting existed.",
				},
			},
			Operations: map[logical.Operation]framework.OperationHandler{
				logical.CreateOperation: &framework.PathOperation{
					Callback: b.roleCreateOrUpdate,
				},
				logical.ReadOperation: &framework.PathOperation{
					Callback: b.roleRead,
				},
				logical.UpdateOperation: &framework.PathOperation{
					Callback: b.roleCreateOrUpdate,
				},
				logical.DeleteOperation: &framework.PathOperation{
					Callback: b.roleDelete,
				},
			},
			ExistenceCheck: b.pathExistenceCheck,
		},
	}
}

func (b *backend) roleCreateOrUpdate(ctx context.Context, req *logical.Request, data *framework.FieldData) (*logical.Response, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	cacheForRatio := data.Get("cache_for_ratio").(int)
	if cacheForRatio <= 0 || cacheForRatio > 100 {
		return logical.ErrorResponse("cache_for_ratio should be greater than 0 and no greater than 100"), nil
	}

	keyType := data.Get("key_type").(string)
	if _, err := getKeyType(keyType); err != nil {
		return logical.ErrorResponse(err.Error()), nil
	}
	if keyType == defaultCertKeyType {
		// Stored as the empty string so that the role hashes, and so caches,
		// exactly as it did before key_type existed: an explicit RSA2048 and
		// no setting at all order the same certificate.
		keyType = ""
	}

	r := role{
		Account:             data.Get("account").(string),
		KeyType:             keyType,
		AllowedDomains:      data.Get("allowed_domains").([]string),
		AllowBareDomains:    data.Get("allow_bare_domains").(bool),
		AllowSubdomains:     data.Get("allow_subdomains").(bool),
		DisableCache:        data.Get("disable_cache").(bool),
		CacheForRatio:       cacheForRatio,
		RevokeOnLeaseExpiry: data.Get("revoke_on_lease_expiry").(bool),
	}
	if err := r.save(ctx, req.Storage, req.Path); err != nil {
		return nil, err
	}

	resp, err := b.roleRead(ctx, req, data)
	if err != nil {
		return nil, err
	}
	if resp != nil && cacheForRatio >= cacheForRatioServesUntilExpiry {
		resp.AddWarning(
			"cache_for_ratio of 100 serves a cached certificate for the whole of its lifetime, so a request " +
				"arriving just before the entry is dropped receives one that has expired or is about to. " +
				"Leave room for whatever renewal window the consumers of this role need.")
	}

	return resp, nil
}

func (b *backend) roleRead(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	r, err := getRole(ctx, req.Storage, req.Path)
	if err != nil {
		return nil, err
	}
	if r == nil {
		return logical.ErrorResponse("This role does not exist"), nil
	}

	return &logical.Response{
		Data: map[string]interface{}{
			"account":                r.Account,
			"allowed_domains":        r.AllowedDomains,
			"allow_bare_domains":     r.AllowBareDomains,
			"allow_subdomains":       r.AllowSubdomains,
			"disable_cache":          r.DisableCache,
			"cache_for_ratio":        r.CacheForRatio,
			"revoke_on_lease_expiry": r.RevokeOnLeaseExpiry,
			"key_type":               r.keyTypeName(),
		},
	}, nil
}

func (b *backend) roleDelete(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	return nil, req.Storage.Delete(ctx, req.Path)
}

func (b *backend) roleList(ctx context.Context, req *logical.Request, _ *framework.FieldData) (*logical.Response, error) {
	entries, err := req.Storage.List(ctx, "roles/")
	if err != nil {
		return nil, err
	}

	return logical.ListResponse(entries), nil
}

type role struct {
	Account          string
	AllowedDomains   []string
	AllowBareDomains bool
	AllowSubdomains  bool
	DisableCache     bool
	CacheForRatio    int
	// Kept out of the JSON encoding because getCacheKey hashes this struct:
	// revocation policy does not change the certificate that gets issued, so
	// toggling it must not orphan every cache entry under the role.
	// mapstructure ignores json tags, so the field still round-trips through
	// storage in save/getRole.
	RevokeOnLeaseExpiry bool `json:"-"`
	// The key type does change the certificate, so it belongs in the hash;
	// omitempty keeps roles that never set it, and roles set to the default,
	// on the cache keys they had before the field existed.
	KeyType string `json:",omitempty"`
}

// defaultCertKeyType is what every certificate was issued with before a role
// could choose, and what an unset key_type still means.
const defaultCertKeyType = "RSA2048"

// keyTypeName returns the role's key type as a caller would have written it.
func (r *role) keyTypeName() string {
	if r.KeyType == "" {
		return defaultCertKeyType
	}
	return r.KeyType
}

func getRole(ctx context.Context, storage logical.Storage, path string) (*role, error) {
	storageEntry, err := storage.Get(ctx, path)
	if err != nil {
		return nil, err
	}
	if storageEntry == nil {
		return nil, nil
	}

	var d map[string]interface{}
	err = storageEntry.DecodeJSON(&d)
	if err != nil {
		return nil, err
	}

	var r *role
	err = mapstructure.Decode(d, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *role) save(ctx context.Context, storage logical.Storage, path string) error {
	var data map[string]interface{}
	err := mapstructure.Decode(r, &data)
	if err != nil {
		return err
	}

	storageEntry, err := logical.StorageEntryJSON(path, data)
	if err != nil {
		return err
	}

	return storage.Put(ctx, storageEntry)
}
