package timestamp

import (
	"go-sdk/interfaces"
	. "go-sdk/metadata"
)

type StorageNowValue = uint64
type StorageNow struct{}

func (this *StorageNow) PalletName() string {
	return PalletName
}

func (this *StorageNow) StorageName() string {
	return "Now"
}

func (this *StorageNow) Fetch(blockStorage interfaces.BlockStorageT) (StorageNowValue, error) {
	val, err := GenericFetch[StorageNowValue](blockStorage, this)
	if err != nil {
		return 0, err
	}

	return val.Unwrap(), nil
}
