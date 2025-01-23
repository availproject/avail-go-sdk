package transaction_payment

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
)

type StorageNextFeeMultiplierValue = Balance
type StorageNextFeeMultiplier struct{}

func (this *StorageNextFeeMultiplier) PalletName() string {
	return PalletName
}

func (this *StorageNextFeeMultiplier) StorageName() string {
	return "NextFeeMultiplier"
}

func (this *StorageNextFeeMultiplier) Fetch(blockStorage interfaces.BlockStorageT) (StorageNextFeeMultiplierValue, error) {
	val, err := GenericFetch[StorageNextFeeMultiplierValue](blockStorage, this)
	if err != nil {
		return StorageNextFeeMultiplierValue{}, err
	}

	return val.Unwrap(), nil
}
