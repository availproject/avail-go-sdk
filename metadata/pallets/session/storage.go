package session

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageCurrentIndexValue = uint32
type StorageCurrentIndex struct{}

func (sci *StorageCurrentIndex) PalletName() string {
	return PalletName
}

func (sci *StorageCurrentIndex) StorageName() string {
	return "CurrentIndex"
}

func (sci *StorageCurrentIndex) Fetch(blockStorage interfaces.BlockStorageT) (StorageCurrentIndexValue, error) {
	return GenericFetchDefault[StorageCurrentIndexValue](blockStorage, sci)
}

//
//
//

type StorageDisabledValidatorsValue = []uint32
type StorageDisabledValidators struct{}

func (sdv *StorageDisabledValidators) PalletName() string {
	return PalletName
}

func (sdv *StorageDisabledValidators) StorageName() string {
	return "DisabledValidators"
}

func (sdv *StorageDisabledValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageDisabledValidatorsValue, error) {
	val, err := GenericFetch[StorageDisabledValidatorsValue](blockStorage, sdv)
	return val.UnwrapOr(StorageDisabledValidatorsValue{}), err
}

//
//
//

type StorageKeyOwnerKey = Tuple2[[4]byte, []byte]
type StorageKeyOwnerValue = prim.AccountId
type StorageKeyOwnerEntry = StorageEntry[StorageKeyOwnerKey, StorageKeyOwnerValue]

type StorageKeyOwner struct{}

func (sko *StorageKeyOwner) PalletName() string {
	return PalletName
}

func (sko *StorageKeyOwner) StorageName() string {
	return "KeyOwner"
}

func (sko *StorageKeyOwner) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sko *StorageKeyOwner) Fetch(blockStorage interfaces.BlockStorageT, key StorageKeyOwnerKey) (prim.Option[StorageKeyOwnerEntry], error) {
	return GenericMapFetch[StorageKeyOwnerValue](blockStorage, key, sko)

}

func (sko *StorageKeyOwner) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageKeyOwnerEntry, error) {
	return GenericMapKeysFetch[StorageKeyOwnerValue, StorageKeyOwnerKey](blockStorage, sko)
}

//
//
//

type StorageNextKeysKey = prim.AccountId
type StorageNextKeysValue = SessionKeys
type StorageNextKeysEntry = StorageEntry[StorageNextKeysKey, StorageNextKeysValue]

type StorageNextKeys struct{}

func (snk *StorageNextKeys) PalletName() string {
	return PalletName
}

func (snk *StorageNextKeys) StorageName() string {
	return "NextKeys"
}

func (snk *StorageNextKeys) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (snk *StorageNextKeys) Fetch(blockStorage interfaces.BlockStorageT, key StorageNextKeysKey) (prim.Option[StorageNextKeysEntry], error) {
	return GenericMapFetch[StorageNextKeysValue](blockStorage, key, snk)

}

func (snk *StorageNextKeys) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageNextKeysEntry, error) {
	return GenericMapKeysFetch[StorageNextKeysValue, StorageNextKeysKey](blockStorage, snk)
}

//
//
//

type StorageQueuedChangedValue = bool
type StorageQueuedChanged struct{}

func (sqc *StorageQueuedChanged) PalletName() string {
	return PalletName
}

func (sqc *StorageQueuedChanged) StorageName() string {
	return "QueuedChanged"
}

func (sqc *StorageQueuedChanged) Fetch(blockStorage interfaces.BlockStorageT) (StorageQueuedChangedValue, error) {
	return GenericFetchDefault[StorageQueuedChangedValue](blockStorage, sqc)
}

//
//
//

type StorageQueuedKeysValue = []Tuple2[prim.AccountId, SessionKeys]
type StorageQueuedKeys struct{}

func (sqk *StorageQueuedKeys) PalletName() string {
	return PalletName
}

func (sqk *StorageQueuedKeys) StorageName() string {
	return "QueuedKeys"
}

func (sqk *StorageQueuedKeys) Fetch(blockStorage interfaces.BlockStorageT) (StorageQueuedKeysValue, error) {
	val, err := GenericFetch[StorageQueuedKeysValue](blockStorage, sqk)
	return val.UnwrapOr(StorageQueuedKeysValue{}), err
}

//
//
//

type StorageValidatorsValue = []prim.AccountId
type StorageValidators struct{}

func (sv *StorageValidators) PalletName() string {
	return PalletName
}

func (sv *StorageValidators) StorageName() string {
	return "Validators"
}

func (sv *StorageValidators) Fetch(blockStorage interfaces.BlockStorageT) (StorageValidatorsValue, error) {
	val, err := GenericFetch[StorageValidatorsValue](blockStorage, sv)
	return val.UnwrapOr(StorageValidatorsValue{}), err
}
