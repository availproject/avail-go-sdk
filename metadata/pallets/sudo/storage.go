package sudo

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
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
