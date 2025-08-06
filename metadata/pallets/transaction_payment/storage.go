package transaction_payment

import (
	"github.com/availproject/avail-go-sdk/interfaces"
	. "github.com/availproject/avail-go-sdk/metadata"
)

type StorageNextFeeMultiplierValue = Balance
type StorageNextFeeMultiplier struct{}

func (snfm *StorageNextFeeMultiplier) PalletName() string {
	return PalletName
}

func (snfm *StorageNextFeeMultiplier) StorageName() string {
	return "NextFeeMultiplier"
}

func (snfm *StorageNextFeeMultiplier) Fetch(blockStorage interfaces.BlockStorageT) (StorageNextFeeMultiplierValue, error) {
	val, err := GenericFetch[StorageNextFeeMultiplierValue](blockStorage, snfm)
	if err != nil {
		return StorageNextFeeMultiplierValue{}, err
	}

	return val.Unwrap(), nil
}
