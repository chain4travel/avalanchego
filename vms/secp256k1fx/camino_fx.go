// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package secp256k1fx

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/hashing"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/multisig"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	errNotSecp256Cred  = errors.New("expected secp256k1 credentials")
	errWrongOutputType = errors.New("wrong output type")
	errNotAliasGetter  = errors.New("state isn't msig alias getter")
	ErrMsigCombination = errors.New("msig combinations not supported")
)

type Owned interface {
	Owners() interface{}
}

type AliasGetter interface {
	GetMultisigAlias(ids.ShortID) (*multisig.AliasWithNonce, error)
}

type (
	RecoverMap map[ids.ShortID][crypto.SECP256K1RSigLen]byte
)

type CaminoFx struct {
	Fx
}

func (fx *CaminoFx) Initialize(vmIntf interface{}) error {
	err := fx.Fx.Initialize(vmIntf)
	if err != nil {
		return err
	}

	c := fx.VM.CodecRegistry()
	if camino, ok := c.(codec.CaminoRegistry); ok {
		if err := camino.RegisterCustomType(&MultisigCredential{}); err != nil {
			return err
		}
	}
	return nil
}

func (fx *CaminoFx) VerifyPermission(txIntf, inIntf, credIntf, ownerIntf interface{}) error {
	if cred, ok := credIntf.(*MultisigCredential); ok {
		credIntf = &cred.Credential
	}
	return fx.Fx.VerifyPermission(txIntf, inIntf, credIntf, ownerIntf)
}

func (fx *CaminoFx) VerifyOperation(txIntf, opIntf, credIntf interface{}, utxosIntf []interface{}) error {
	if cred, ok := credIntf.(*MultisigCredential); ok {
		credIntf = &cred.Credential
	}
	return fx.Fx.VerifyOperation(txIntf, opIntf, credIntf, utxosIntf)
}

func (fx *CaminoFx) VerifyTransfer(txIntf, inIntf, credIntf, utxoIntf interface{}) error {
	if cred, ok := credIntf.(*MultisigCredential); ok {
		credIntf = &cred.Credential
	}
	return fx.Fx.VerifyTransfer(txIntf, inIntf, credIntf, utxoIntf)
}

func (fx *Fx) RecoverAddresses(utx UnsignedTx, verifies []verify.Verifiable) (RecoverMap, error) {
	ret := make(RecoverMap, len(verifies))
	visited := make(map[[crypto.SECP256K1RSigLen]byte]bool)

	txHash := hashing.ComputeHash256(utx.Bytes())
	for _, v := range verifies {
		cred, ok := v.(CredentialIntf)
		if !ok {
			return nil, errNotSecp256Cred
		}
		for _, sig := range cred.Signatures() {
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
	out, ok := outIntf.(Owned)
	if !ok {
		return errWrongOutputType
	}
	owners, ok := out.Owners().(*OutputOwners)
	if !ok {
		return errWrongOwnerType
	}
	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	// We don't support msig combinations / nesting for now
	if len(owners.Addrs) == 1 {
		alias, err := msig.GetMultisigAlias(owners.Addrs[0])
		if err == database.ErrNotFound {
			return nil
		} else if err != nil && err != database.ErrNotFound {
			return err
		}
		aliasOwners, ok := alias.Owners.(*OutputOwners)
		if !ok {
			return errWrongOwnerType
		}
		owners = aliasOwners
	}

	for _, addr := range owners.Addrs {
		if _, err := msig.GetMultisigAlias(addr); err != nil && err != database.ErrNotFound {
			return err
		} else if err == nil {
			return ErrMsigCombination
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
	cred, ok := credIntf.(CredentialIntf)
	if !ok {
		return errWrongCredentialType
	}
	out, ok := utxoIntf.(TransferOutputIntf)
	if !ok {
		return errWrongUTXOType
	}
	owners, ok := out.Owners().(*OutputOwners)
	if !ok {
		return errWrongOwnerType
	}
	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return errNotAliasGetter
	}

	if err := verify.All(out, in, cred); err != nil {
		return err
	} else if out.Amount() != in.Amt {
		return fmt.Errorf("out amount and input differ")
	}

	return fx.verifyMultisigCredentials(tx, &in.Input, cred, owners, msig)
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
	cred, ok := credIntf.(CredentialIntf)
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

func (*Fx) CollectMultisigAliases(ownerIntf, msigIntf interface{}) ([]interface{}, error) {
	owners, ok := ownerIntf.(*OutputOwners)
	if !ok {
		return nil, errWrongUTXOType
	}
	msig, ok := msigIntf.(AliasGetter)
	if !ok {
		return nil, errNotAliasGetter
	}

	return collectMultisigAliases(owners, msig)
}

func (fx *Fx) verifyMultisigCredentials(tx UnsignedTx, in *Input, cred CredentialIntf, owners *OutputOwners, msig AliasGetter) error {
	sigIdxs := cred.SignatureIndices()
	if sigIdxs == nil {
		sigIdxs = in.SigIndices
	}

	if len(sigIdxs) != len(cred.Signatures()) {
		return errInputCredentialSignersMismatch
	}

	resolved, err := fx.RecoverAddresses(tx, []verify.Verifiable{cred})
	if err != nil {
		return err
	}

	tf := func(
		addr ids.ShortID,
		totalVisited,
		totalVerified uint32,
	) (bool, error) {
		// check that input sig index matches
		if totalVerified >= uint32(len(sigIdxs)) {
			return false, nil
		}

		if sig, exists := resolved[addr]; exists &&
			sig == cred.Signatures()[totalVerified] &&
			sigIdxs[totalVerified] == totalVisited {
			return true, nil
		}
		return false, nil
	}

	sigsVerified, err := TraverseOwners(owners, msig, tf)
	if err != nil {
		return err
	}
	if sigsVerified < uint32(len(sigIdxs)) {
		return errTooManySigners
	}

	return nil
}

func collectMultisigAliases(owners *OutputOwners, msig AliasGetter) ([]interface{}, error) {
	result := make([]interface{}, 0, len(owners.Addrs))

	tf := func(alias *multisig.AliasWithNonce) {
		result = append(result, alias)
	}

	if err := TraverseAliases(owners, msig, tf); err != nil {
		return nil, err
	}
	return result, nil
}

// ExtractFromAndSigners splits an array of PrivateKeys into `from` and `signers`
// The delimiter is a `nil` PrivateKey.
// If no delimiter exists, the given PrivateKeys are used for both from and signing
// Having different sets of `from` and `signer` allows multisignature feature
func ExtractFromAndSigners(keys []*crypto.PrivateKeySECP256K1R) (set.Set[ids.ShortID], []*crypto.PrivateKeySECP256K1R) {
	// find nil key which splits froms and signer
	splitIndex := len(keys)
	for index, key := range keys {
		if key == nil {
			splitIndex = index
			break
		}
	}

	if splitIndex == len(keys) {
		from := set.NewSet[ids.ShortID](len(keys))
		for _, key := range keys {
			from.Add(key.PublicKey().Address())
		}
		return from, keys
	}

	// Addresses we get UTXOs for
	from := set.NewSet[ids.ShortID](splitIndex)
	for index := 0; index < splitIndex; index++ {
		from.Add(keys[index].PublicKey().Address())
	}

	// Signers which will signe the Outputs
	signer := make([]*crypto.PrivateKeySECP256K1R, len(keys)-splitIndex-1)
	for index := 0; index < len(signer); index++ {
		signer[index] = keys[index+splitIndex+1]
	}
	return from, signer
}
