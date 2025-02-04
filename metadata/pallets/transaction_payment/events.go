package transaction_payment

import (
	"github.com/availproject/avail-go-sdk/primitives"

	"github.com/itering/scale.go/utiles/uint128"
)

// A transaction fee `actual_fee`, of which `tip` was added to the minimum inclusion fee, has been paid by `who`.
type EventTransactionFeePaid struct {
	Who       primitives.AccountId
	ActualFee uint128.Uint128
	Tip       uint128.Uint128
}

func (this EventTransactionFeePaid) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventTransactionFeePaid) PalletName() string {
	return PalletName
}

func (this EventTransactionFeePaid) EventIndex() uint8 {
	return 0
}

func (this EventTransactionFeePaid) EventName() string {
	return "TransactionFeePaid"
}
