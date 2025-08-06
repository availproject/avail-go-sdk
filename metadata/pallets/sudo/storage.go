package sudo

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageKeyValue = prim.AccountId
type StorageKey struct{}

func (sk *StorageKey) PalletName() string {
	return PalletName
}

func (sk *StorageKey) StorageName() string {
	return "Key"
}

func (sk *StorageKey) Fetch(blockStorage interfaces.BlockStorageT) (prim.Option[StorageKeyValue], error) {
	return GenericFetch[StorageKeyValue](blockStorage, sk)
}
