package identity

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageIdentityOfKey = prim.AccountId
type StorageIdentityOfValue = Tuple2[Registration, prim.Option[[]byte]]
type StorageIdentityOfEntry = StorageEntry[StorageIdentityOfKey, StorageIdentityOfValue]

type StorageIdentityOf struct{}

func (sio *StorageIdentityOf) PalletName() string {
	return PalletName
}

func (sio *StorageIdentityOf) StorageName() string {
	return "IdentityOf"
}

func (sio *StorageIdentityOf) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (sio *StorageIdentityOf) Fetch(blockStorage interfaces.BlockStorageT, key StorageIdentityOfKey) (prim.Option[StorageIdentityOfEntry], error) {
	return GenericMapFetch[StorageIdentityOfValue](blockStorage, key, sio)
}

func (sio *StorageIdentityOf) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageIdentityOfEntry, error) {
	return GenericMapKeysFetch[StorageIdentityOfValue, StorageIdentityOfKey](blockStorage, sio)
}
