package data_availability

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
)

type StorageNextAppId struct {
	Value uint32 `scale:"compact"`
}

func (this *StorageNextAppId) PalletName() string {
	return PalletName
}

func (this *StorageNextAppId) StorageName() string {
	return "NextAppId"
}

func (this *StorageNextAppId) Fetch(blockStorage interfaces.BlockStorageT) (uint32, error) {
	val, err := GenericFetch[StorageNextAppId](blockStorage, this)
	if err != nil {
		return 0, err
	}

	return val.Unwrap().Value, nil
}

//
//
//

type StorageAppKeysKey = []byte
type StorageAppKeysEntry = StorageEntry[StorageAppKeysKey, StorageAppKeys]

type StorageAppKeys struct {
	Owner AccountId
	AppId uint32 `scale:"compact"`
}

func (this *StorageAppKeys) PalletName() string {
	return PalletName
}

func (this *StorageAppKeys) StorageName() string {
	return "AppKeys"
}

func (this *StorageAppKeys) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (this *StorageAppKeys) Fetch(blockStorage interfaces.BlockStorageT, key StorageAppKeysKey) (prim.Option[StorageAppKeysEntry], error) {
	return GenericMapFetch[StorageAppKeys](blockStorage, key, this)
}

func (this *StorageAppKeys) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageAppKeysEntry, error) {
	return GenericMapKeysFetch[StorageAppKeys, StorageAppKeysKey](blockStorage, this)
}
