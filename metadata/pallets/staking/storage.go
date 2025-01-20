package staking

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
)

type StorageActiveEra struct {
	Index uint32
	Start prim.Option[uint64]
}

func (this *StorageActiveEra) PalletName() string {
	return PalletName
}

func (this *StorageActiveEra) StorageName() string {
	return "ActiveEra"
}

func (this *StorageActiveEra) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageActiveEra], error) {
	return GenericFetch[StorageActiveEra](blockStorage, this)
}

//
//
//

type StorageBondedKey = AccountId
type StorageBondedValue = AccountId
type StorageBondedEntry = StorageEntry[StorageBondedKey, StorageBondedValue]

type StorageBonded struct{}

func (this *StorageBonded) PalletName() string {
	return "Staking"
}

func (this *StorageBonded) StorageName() string {
	return "Bonded"
}

func (this *StorageBonded) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageBonded) Fetch(blockStorage interfaces.BlockStorageT, key StorageBondedKey) (prim.Option[StorageBondedEntry], error) {
	return GenericMapFetch[StorageBondedValue](blockStorage, key, this)
}

func (this *StorageBonded) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageBondedEntry, error) {
	return GenericFetchKeys[StorageBondedValue, StorageBondedKey](blockStorage, this)
}

//
//
//

type BondedEra struct {
	Tup1 uint32
	Tup2 uint32
}

type BondedErasValue = []BondedEra
type StorageBondedEras struct{}

func (this *StorageBondedEras) PalletName() string {
	return PalletName
}

func (this *StorageBondedEras) StorageName() string {
	return "BondedEras"
}

func (this *StorageBondedEras) Fetch(blockStorage interfaces.BlockStorageT) (BondedErasValue, error) {
	val, err := GenericFetch[BondedErasValue](blockStorage, this)
	if err != nil {
		return nil, err
	}

	if val.IsNone() {
		return BondedErasValue{}, nil
	}

	return val.Unwrap(), nil
}

//
//
//

type StorageCanceledSlashPayoutValue = Balance
type StorageCanceledSlashPayout struct{}

func (this *StorageCanceledSlashPayout) PalletName() string {
	return PalletName
}

func (this *StorageCanceledSlashPayout) StorageName() string {
	return "CanceledSlashPayout"
}

func (this *StorageCanceledSlashPayout) Fetch(blockStorage interfaces.BlockStorageT) (StorageCanceledSlashPayoutValue, error) {
	val, err := GenericFetch[StorageCanceledSlashPayoutValue](blockStorage, this)
	if err != nil {
		return StorageCanceledSlashPayoutValue{}, err
	}

	return val.Unwrap(), nil
}

//
//
//

type StorageChillThresholdValue = uint8
type StorageChillThreshold struct{}

func (this *StorageChillThreshold) PalletName() string {
	return PalletName
}

func (this *StorageChillThreshold) StorageName() string {
	return "CanceledSlashPayout"
}

func (this *StorageChillThreshold) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageChillThresholdValue], error) {
	val, err := GenericFetch[StorageChillThresholdValue](blockStorage, this)
	if err != nil {
		return prim.NewNone[StorageChillThresholdValue](), err
	}

	return val, nil
}
