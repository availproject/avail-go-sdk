package balances

import (
	metadata "github.com/availproject/avail-go-sdk/metadata"
)

// Do not add, remove or change any of the field members.
type EventDeposit struct {
	Who    metadata.AccountId
	Amount metadata.Balance
}

func (this EventDeposit) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventDeposit) PalletName() string {
	return PalletName
}

func (this EventDeposit) EventIndex() uint8 {
	return 7
}

func (this EventDeposit) EventName() string {
	return "Deposit"
}

// Do not add, remove or change any of the field members.
type EventWithdraw struct {
	Who    metadata.AccountId
	Amount metadata.Balance
}

func (this EventWithdraw) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventWithdraw) PalletName() string {
	return PalletName
}

func (this EventWithdraw) EventIndex() uint8 {
	return 8
}

func (this EventWithdraw) EventName() string {
	return "Withdraw"
}
