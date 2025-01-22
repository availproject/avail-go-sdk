package sudo

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
	prim "go-sdk/primitives"
)

type StorageKeyValue = AccountId
type StorageKey struct{}

func (this *StorageKey) PalletName() string {
	return PalletName
}

func (this *StorageKey) StorageName() string {
	return "Key"
}

func (this *StorageKey) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageKeyValue], error) {
	return GenericFetch[StorageKeyValue](blockStorage, this)
}