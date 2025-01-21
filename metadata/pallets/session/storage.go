package session

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
)

type StorageCurrentIndexValue = uint32
type StorageCurrentIndex struct{}

func (this *StorageCurrentIndex) PalletName() string {
	return PalletName
}

func (this *StorageCurrentIndex) StorageName() string {
	return "CurrentIndex"
}

func (this *StorageCurrentIndex) Fetch(blockStorage interfaces.BlockStorageT) (StorageCurrentIndexValue, error) {
	return GenericFetchDefault[StorageCurrentIndexValue](blockStorage, this)
}

//
//
//

type StorageDisabledValidatorsValue = []uint32
type StorageDisabledValidators struct{}

func (this *StorageDisabledValidators) PalletName() string {
	return PalletName
}

func (this *StorageDisabledValidators) StorageName() string {
	return "DisabledValidators"
}

func (this *StorageDisabledValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageDisabledValidatorsValue, error) {
	val, err := GenericFetch[StorageDisabledValidatorsValue](blockStorage, this)
	return val.UnwrapOr(StorageDisabledValidatorsValue{}), err
}

//
//
//

type StorageKeyOwnerKey = Tuple2[[4]byte, []byte]
type StorageKeyOwnerValue = AccountId
type StorageKeyOwnerEntry = StorageEntry[StorageKeyOwnerKey, StorageKeyOwnerValue]

type StorageKeyOwner struct{}

func (this *StorageKeyOwner) PalletName() string {
	return PalletName
}

func (this *StorageKeyOwner) StorageName() string {
	return "KeyOwner"
}

func (this *StorageKeyOwner) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageKeyOwner) Fetch(blockStorage interfaces.BlockStorageT, key StorageKeyOwnerKey) (prim.Option[StorageKeyOwnerEntry], error) {
	return GenericMapFetch[StorageKeyOwnerValue](blockStorage, key, this)

}

func (this *StorageKeyOwner) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageKeyOwnerEntry, error) {
	return GenericMapKeysFetch[StorageKeyOwnerValue, StorageKeyOwnerKey](blockStorage, this)
}

//
//
//

type StorageNextKeysKey = AccountId
type StorageNextKeysValue = SessionKeys
type StorageNextKeysEntry = StorageEntry[StorageNextKeysKey, StorageNextKeysValue]

type StorageNextKeys struct{}

func (this *StorageNextKeys) PalletName() string {
	return PalletName
}

func (this *StorageNextKeys) StorageName() string {
	return "NextKeys"
}

func (this *StorageNextKeys) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageNextKeys) Fetch(blockStorage interfaces.BlockStorageT, key StorageNextKeysKey) (prim.Option[StorageNextKeysEntry], error) {
	return GenericMapFetch[StorageNextKeysValue](blockStorage, key, this)

}

func (this *StorageNextKeys) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageNextKeysEntry, error) {
	return GenericMapKeysFetch[StorageNextKeysValue, StorageNextKeysKey](blockStorage, this)
}

//
//
//

type StorageQueuedChangedValue = bool
type StorageQueuedChanged struct{}

func (this *StorageQueuedChanged) PalletName() string {
	return PalletName
}

func (this *StorageQueuedChanged) StorageName() string {
	return "QueuedChanged"
}

func (this *StorageQueuedChanged) Fetch(blockStorage interfaces.BlockStorageT) (StorageQueuedChangedValue, error) {
	return GenericFetchDefault[StorageQueuedChangedValue](blockStorage, this)
}

//
//
//

type StorageQueuedKeysValue = []Tuple2[AccountId, SessionKeys]
type StorageQueuedKeys struct{}

func (this *StorageQueuedKeys) PalletName() string {
	return PalletName
}

func (this *StorageQueuedKeys) StorageName() string {
	return "QueuedKeys"
}

func (this *StorageQueuedKeys) Fetch(blockStorage interfaces.BlockStorageT) (StorageQueuedKeysValue, error) {
	val, err := GenericFetch[StorageQueuedKeysValue](blockStorage, this)
	return val.UnwrapOr(StorageQueuedKeysValue{}), err
}

//
//
//

type StorageValidatorsValue = []AccountId
type StorageValidators struct{}

func (this *StorageValidators) PalletName() string {
	return PalletName
}

func (this *StorageValidators) StorageName() string {
	return "Validators"
}

func (this *StorageValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageValidatorsValue, error) {
	val, err := GenericFetch[StorageValidatorsValue](blockStorage, this)
	return val.UnwrapOr(StorageValidatorsValue{}), err
}
