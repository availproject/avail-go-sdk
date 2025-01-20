package treasury_committee

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
)

type StorageMembersValue = []AccountId
type StorageMembers struct{}

func (this *StorageMembers) PalletName() string {
	return PalletName
}

func (this *StorageMembers) StorageName() string {
	return "Members"
}

func (this *StorageMembers) Fetch(blockStorage interfaces.BlockStorageT) (StorageMembersValue, error) {
	val, err := GenericFetch[StorageMembersValue](blockStorage, this)
	if err != nil {
		return nil, err
	}

	return val.UnwrapOr(StorageMembersValue{}), nil
}

//
//
//

type StoragePrimeValue = AccountId
type StoragePrime struct{}

func (this *StoragePrime) PalletName() string {
	return PalletName
}

func (this *StoragePrime) StorageName() string {
	return "Prime"
}

func (this *StoragePrime) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StoragePrimeValue], error) {
	return GenericFetch[StoragePrimeValue](blockStorage, this)
}

//
//
//

type StorageProposalCountValue = uint32
type StorageProposalCount struct{}

func (this *StorageProposalCount) PalletName() string {
	return PalletName
}

func (this *StorageProposalCount) StorageName() string {
	return "ProposalCount"
}

func (this *StorageProposalCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageProposalCountValue, error) {
	val, err := GenericFetch[StorageProposalCountValue](blockStorage, this)
	if err != nil {
		return 0, err
	}

	// 0 is default
	return val.Unwrap(), nil
}

//
//
//

type StorageProposalsValue = []prim.H256
type StorageProposals struct{}

func (this *StorageProposals) PalletName() string {
	return PalletName
}

func (this *StorageProposals) StorageName() string {
	return "Proposals"
}

func (this *StorageProposals) Fetch(blockStorage interfaces.BlockStorageT) (StorageProposalsValue, error) {
	val, err := GenericFetch[StorageProposalsValue](blockStorage, this)
	if err != nil {
		return nil, err
	}

	return val.UnwrapOr(StorageProposalsValue{}), nil
}
