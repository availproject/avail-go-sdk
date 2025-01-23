package balances

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageTotalIssuance struct {
	Value Balance
}

func (this *StorageTotalIssuance) PalletName() string {
	return PalletName
}

func (this *StorageTotalIssuance) StorageName() string {
	return "TotalIssuance"
}

func (this *StorageTotalIssuance) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageTotalIssuance], error) {
	return GenericFetch[StorageTotalIssuance](blockStorage, this)
}

//
//
//

type StorageInactiveIssuance struct {
	Value Balance
}

func (this *StorageInactiveIssuance) PalletName() string {
	return PalletName
}

func (this *StorageInactiveIssuance) StorageName() string {
	return "InactiveIssuance"
}

func (this *StorageInactiveIssuance) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageInactiveIssuance], error) {
	return GenericFetch[StorageInactiveIssuance](blockStorage, this)
}

//
//
//

type StorageLocksKey = AccountId
type StorageLocksEntry = StorageEntry[StorageLocksKey, StorageLocks]

type StorageLocks struct {
	Value []BalanceLock
}

func (this *StorageLocks) PalletName() string {
	return PalletName
}

func (this *StorageLocks) StorageName() string {
	return "Locks"
}

func (this *StorageLocks) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (this *StorageLocks) Fetch(blockStorage interfaces.BlockStorageT, key StorageLocksKey) (prim.Option[StorageLocksEntry], error) {
	return GenericMapFetch[StorageLocks](blockStorage, key, this)
}

func (this *StorageLocks) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageLocksEntry, error) {
	return GenericMapKeysFetch[StorageLocks, StorageLocksKey](blockStorage, this)
}

type BalanceLock struct {
	Id      [8]uint8
	Amount  Balance
	Reasons Reasons
}

type Reasons struct {
	VariantIndex uint8
}

func (this Reasons) ToString() string {
	switch this.VariantIndex {
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
