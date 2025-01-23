package identity

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type StorageIdentityOfKey = AccountId
type StorageIdentityOfValue = Tuple2[Registration, prim.Option[[]byte]]
type StorageIdentityOfEntry = StorageEntry[StorageIdentityOfKey, StorageIdentityOfValue]

type StorageIdentityOf struct{}

func (this *StorageIdentityOf) PalletName() string {
	return PalletName
}

func (this *StorageIdentityOf) StorageName() string {
	return "IdentityOf"
}

func (this *StorageIdentityOf) MapKeyHasher() uint8 {
	return Twox64ConcatHasher
}

func (this *StorageIdentityOf) Fetch(blockStorage interfaces.BlockStorageT, key StorageIdentityOfKey) (prim.Option[StorageIdentityOfEntry], error) {
	return GenericMapFetch[StorageIdentityOfValue](blockStorage, key, this)
}

func (this *StorageIdentityOf) FetchAll(blockStorage interfaces.BlockStorageT) ([]StorageIdentityOfEntry, error) {
	return GenericMapKeysFetch[StorageIdentityOfValue, StorageIdentityOfKey](blockStorage, this)
}
