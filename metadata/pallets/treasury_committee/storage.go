package treasury_committee

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageMembersValue = []prim.AccountId
type StorageMembers struct{}

func (sm *StorageMembers) PalletName() string {
	return PalletName
}

func (sm *StorageMembers) StorageName() string {
	return "Members"
}

func (sm *StorageMembers) Fetch(blockStorage interfaces.BlockStorageT) (StorageMembersValue, error) {
	val, err := GenericFetch[StorageMembersValue](blockStorage, sm)
	return val.UnwrapOr(StorageMembersValue{}), err
}

//
//
//

type StoragePrimeValue = prim.AccountId
type StoragePrime struct{}

func (sp *StoragePrime) PalletName() string {
	return PalletName
}

func (sp *StoragePrime) StorageName() string {
	return "Prime"
}

func (sp *StoragePrime) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StoragePrimeValue], error) {
	return GenericFetch[StoragePrimeValue](blockStorage, sp)
}

//
//
//

type StorageProposalCountValue = uint32
type StorageProposalCount struct{}

func (spc *StorageProposalCount) PalletName() string {
	return PalletName
}

func (spc *StorageProposalCount) StorageName() string {
	return "ProposalCount"
}

func (spc *StorageProposalCount) Fetch(blockStorage interfaces.BlockStorageT) (StorageProposalCountValue, error) {
	return GenericFetchDefault[StorageProposalCountValue](blockStorage, spc)
}

//
//
//

type StorageProposalsValue = []prim.H256
type StorageProposals struct{}

func (sp *StorageProposals) PalletName() string {
	return PalletName
}

func (sp *StorageProposals) StorageName() string {
	return "Proposals"
}

func (sp *StorageProposals) Fetch(blockStorage interfaces.BlockStorageT) (StorageProposalsValue, error) {
	return GenericFetchDefault[StorageProposalsValue](blockStorage, sp)
}
