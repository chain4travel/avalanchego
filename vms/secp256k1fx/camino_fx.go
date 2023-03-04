// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"
	"fmt"
	"math"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/vms/components/multisig"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	errNotSecp256Cred    = errors.New("expected secp256k1 credentials")
	errWrongOutputType   = errors.New("wrong output type")
	errMsigCombination   = errors.New("msig combinations not supported")
	errNotAliasGetter    = errors.New("state isn't msig alias getter")
	errTooFewSigIndices  = errors.New("too few signature indices")
	errTooManySigIndices = errors.New("too many signature indices")
)

type Owned interface {
	Owners() interface{}
}

type AliasGetter interface {
	GetMultisigAlias(ids.ShortID) (*multisig.Alias, error)
}

type (
	RecoverMap map[ids.ShortID][crypto.SECP256K1RSigLen]byte
)

func (fx *Fx) RecoverAddresses(utx UnsignedTx, verifies []verify.Verifiable) (RecoverMap, error) {
	ret := make(RecoverMap, len(verifies))
	visited := make(map[[crypto.SECP256K1RSigLen]byte]bool)

	txHash := hashing.ComputeHash256(utx.Bytes())
	for _, v := range verifies {
		cred, ok := v.(*Credential)
		if !ok {
			return nil, errNotSecp256Cred
		}
		for _, sig := range cred.Sigs {
			if visited[sig] {
				continue
			}
			pk, err := fx.SECPFactory.RecoverHashPublicKey(txHash, sig[:])
			if err != nil {
				return nil, err
			}
			visited[sig] = true
			ret[pk.Address()] = sig
		}
	}
	return ret, nil
}

func (*Fx) VerifyMultisigOwner(outIntf, msigIntf interface{}) error {
	out, ok := outIntf.(*TransferOutput)
	if !ok {
		return errWrongOutputType
	}
	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	// We don't support msig combinations / nesting
	if len(out.OutputOwners.Addrs) > 1 {
		for _, addr := range out.OutputOwners.Addrs {
			if _, err := msig.GetMultisigAlias(addr); err != database.ErrNotFound {
				return errMsigCombination
			}
		}
	}

	return nil
}

func (fx *Fx) VerifyMultisigTransfer(txIntf, inIntf, credIntf, utxoIntf, msigIntf interface{}) error {
	tx, ok := txIntf.(UnsignedTx)
	if !ok {
		return errWrongTxType
	}
	in, ok := inIntf.(*TransferInput)
	if !ok {
		return errWrongInputType
	}
	cred, ok := credIntf.(*Credential)
	if !ok {
		return errWrongCredentialType
	}
	out, ok := utxoIntf.(*TransferOutput)
	if !ok {
		return errWrongUTXOType
	}

	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	if err := verify.All(out, in, cred); err != nil {
		return err
	} else if out.Amt != in.Amt {
		return fmt.Errorf("out amount and input differ")
	}

	return fx.verifyMultisigCredentials(tx, &in.Input, cred, &out.OutputOwners, msig)
}

func (fx *Fx) VerifyMultisigPermission(txIntf, inIntf, credIntf, ownerIntf, msigIntf interface{}) error {
	tx, ok := txIntf.(UnsignedTx)
	if !ok {
		return errWrongTxType
	}
	in, ok := inIntf.(*Input)
	if !ok {
		return errWrongInputType
	}
	cred, ok := credIntf.(*Credential)
	if !ok {
		return errWrongCredentialType
	}
	owners, ok := ownerIntf.(*OutputOwners)
	if !ok {
		return errWrongUTXOType
	}

	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	if err := verify.All(owners, in, cred); err != nil {
		return err
	}

	return fx.verifyMultisigCredentials(tx, in, cred, owners, msig)
}

func (fx *Fx) VerifyMultisigUnorderedPermission(txIntf, credIntf, ownerIntf, msigIntf interface{}) error {
	tx, ok := txIntf.(UnsignedTx)
	if !ok {
		return errWrongTxType
	}
	cred, ok := credIntf.(*Credential)
	if !ok {
		return errWrongCredentialType
	}
	owners, ok := ownerIntf.(*OutputOwners)
	if !ok {
		return errWrongUTXOType
	}

	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	if err := verify.All(owners, cred); err != nil {
		return err
	}

	return fx.verifyMultisigUnorderedCredentials(tx, cred, owners, msig)
}

func (fx *Fx) verifyMultisigCredentials(tx UnsignedTx, in *Input, cred *Credential, owners *OutputOwners, msig AliasGetter) error {
	if len(in.SigIndices) > int(owners.Threshold) {
		return errTooManySigIndices
	} else if len(in.SigIndices) < int(owners.Threshold) {
		return errTooFewSigIndices
	}

	resolved, err := fx.RecoverAddresses(tx, []verify.Verifiable{cred})
	if err != nil {
		return err
	}

	tf := func(
		alias bool,
		addr ids.ShortID,
		depth int,
		visited,
		verified,
		totalVisited uint32,
	) (bool, error) {
		if alias {
			return false, nil
		}
		// check that input sig index matches
		if totalVisited >= uint32(len(cred.Sigs)) {
			return false, errTooFewSigIndices
		}

		if sig, exists := resolved[addr]; exists &&
			sig == cred.Sigs[totalVisited] &&
			(depth > 1 || (in.SigIndices[verified] == math.MaxUint32 ||
				in.SigIndices[verified] == visited)) {
			return true, nil
		}
		return false, nil
	}

	var sigsVerified uint32
	if sigsVerified, err = TraverseOwners(owners, msig, tf); err != nil {
		return err
	}

	if sigsVerified < uint32(len(cred.Sigs)) {
		return errTooManySigners
	}

	return nil
}

func (fx *Fx) verifyMultisigUnorderedCredentials(tx UnsignedTx, cred *Credential, owners *OutputOwners, msig AliasGetter) error {
	resolved, err := fx.RecoverAddresses(tx, []verify.Verifiable{cred})
	if err != nil {
		return err
	}

	tf := func(alias bool, addr ids.ShortID, _ int, _, _, _ uint32) (bool, error) {
		if alias {
			return false, nil
		}
		if _, exists := resolved[addr]; exists {
			return true, nil
		}
		return false, nil
	}

	if _, err = TraverseOwners(owners, msig, tf); err != nil {
		return err
	}

	return nil
}
