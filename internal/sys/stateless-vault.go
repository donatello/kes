// Copyright 2022 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package sys

import (
	"context"
	"net/http"

	"github.com/minio/kes"
	"github.com/minio/kes/internal/auth"
	"github.com/minio/kes/internal/key"
)

// NewStatelessVault returns a new Vault with a single Enclave
// that uses the given key store, policy set and identity set.
//
// The Vault is not able to create or delete enclaves.
func NewStatelessVault(operator kes.Identity, keys key.Store, policies auth.PolicySet, identites auth.IdentitySet) Vault {
	return &statelessVault{
		enclave: &Enclave{
			keys:       keys,
			policies:   policies,
			identities: identites,
		},
		operator: operator,
	}
}

type statelessVault struct {
	enclave  *Enclave
	operator kes.Identity
}

var _ Vault = (*statelessVault)(nil) // compiler check

func (v *statelessVault) Seal(ctx context.Context) error { return nil }

func (v *statelessVault) Unseal(context.Context) error { return nil }

func (v *statelessVault) Operator(_ context.Context) (kes.Identity, error) {
	return v.operator, nil
}

func (v *statelessVault) CreateEnclave(_ context.Context, _ string) (*Enclave, error) {
	return nil, kes.NewError(http.StatusNotImplemented, "creating encalves is not supported")
}

func (v *statelessVault) GetEnclave(_ context.Context, name string) (*Enclave, error) {
	if name == "" {
		return v.enclave, nil
	}
	return nil, kes.ErrEnclaveNotFound
}

func (v *statelessVault) DeleteEnclave(_ context.Context, _ string) error {
	return kes.NewError(http.StatusNotImplemented, "deleting encalves is not supported")
}
