package timestamp

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
)

type StorageNowValue = uint64
type StorageNow struct{}

func (sn *StorageNow) PalletName() string {
	return PalletName
}

func (sn *StorageNow) StorageName() string {
	return "Now"
}

func (sn *StorageNow) Fetch(blockStorage interfaces.BlockStorageT) (StorageNowValue, error) {
	val, err := GenericFetch[StorageNowValue](blockStorage, sn)
	if err != nil {
		return 0, err
	}

	return val.Unwrap(), nil
}
