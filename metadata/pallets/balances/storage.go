package balances

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageTotalIssuance struct {
	Value Balance
}

func (sti *StorageTotalIssuance) PalletName() string {
	return PalletName
}

func (sti *StorageTotalIssuance) StorageName() string {
	return "TotalIssuance"
}

func (sti *StorageTotalIssuance) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageTotalIssuance], error) {
	return GenericFetch[StorageTotalIssuance](blockStorage, sti)
}

//
//
//

type StorageInactiveIssuance struct {
	Value Balance
}

func (sii *StorageInactiveIssuance) PalletName() string {
	return PalletName
}

func (sii *StorageInactiveIssuance) StorageName() string {
	return "InactiveIssuance"
}

func (sii *StorageInactiveIssuance) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageInactiveIssuance], error) {
	return GenericFetch[StorageInactiveIssuance](blockStorage, sii)
}

//
//
//

type StorageLocksKey = prim.AccountId
type StorageLocksEntry = StorageEntry[StorageLocksKey, StorageLocks]

type StorageLocks struct {
	Value []BalanceLock
}

func (sl *StorageLocks) PalletName() string {
	return PalletName
}

func (sl *StorageLocks) StorageName() string {
	return "Locks"
}

func (sl *StorageLocks) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (sl *StorageLocks) Fetch(blockStorage interfaces.BlockStorageT, key StorageLocksKey) (prim.Option[StorageLocksEntry], error) {
	return GenericMapFetch[StorageLocks](blockStorage, key, sl)
}

func (sl *StorageLocks) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageLocksEntry, error) {
	return GenericMapKeysFetch[StorageLocks, StorageLocksKey](blockStorage, sl)
}

type BalanceLock struct {
	Id      [8]uint8
	Amount  Balance
	Reasons Reasons
}

type Reasons struct {
	VariantIndex uint8
}

func (r Reasons) ToString() string {
	switch r.VariantIndex {
	case 0:
		return "Fee"
	case 1:
		return "Misc"
	case 2:
		return "All"
	default:
		panic("Unknown Reasons Variant Index")
	}
}
