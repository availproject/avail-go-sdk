package data_availability

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageNextAppId struct {
	Value uint32 `scale:"compact"`
}

func (snai *StorageNextAppId) PalletName() string {
	return PalletName
}

func (snai *StorageNextAppId) StorageName() string {
	return "NextAppId"
}

func (snai *StorageNextAppId) Fetch(blockStorage interfaces.BlockStorageT) (uint32, error) {
	val, err := GenericFetch[StorageNextAppId](blockStorage, snai)
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
	Owner prim.AccountId
	AppId uint32 `scale:"compact"`
}

func (sak *StorageAppKeys) PalletName() string {
	return PalletName
}

func (sak *StorageAppKeys) StorageName() string {
	return "AppKeys"
}

func (sak *StorageAppKeys) MapKeyHasher() uint8 {
	return Blake2_128ConcatHasher
}

func (sak *StorageAppKeys) Fetch(blockStorage interfaces.BlockStorageT, key StorageAppKeysKey) (prim.Option[StorageAppKeysEntry], error) {
	return GenericMapFetch[StorageAppKeys](blockStorage, key, sak)
}

func (sak *StorageAppKeys) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageAppKeysEntry, error) {
	return GenericMapKeysFetch[StorageAppKeys, StorageAppKeysKey](blockStorage, sak)
}
