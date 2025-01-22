package balances

import (
	"go-sdk/metadata"
)

// An account was created with some free balance.
type EventEndowed struct {
	Account     metadata.AccountId
	FreeBalance metadata.Balance
}

func (this EventEndowed) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventEndowed) PalletName() string {
	return PalletName
}

func (this EventEndowed) EventIndex() uint8 {
	return 0
}

func (this EventEndowed) EventName() string {
	return "Endowed"
}

// An account was removed whose balance was non-zero but below ExistentialDeposit, resulting in an outright loss.
type EventDustLost struct {
	Account metadata.AccountId
	Amount  metadata.Balance
}

func (this EventDustLost) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventDustLost) PalletName() string {
	return PalletName
}

func (this EventDustLost) EventIndex() uint8 {
	return 1
}

func (this EventDustLost) EventName() string {
	return "DustLost"
}

// Transfer succeeded.
type EventTransfer struct {
	From   metadata.AccountId
	To     metadata.AccountId
	Amount metadata.Balance
}

func (this EventTransfer) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventTransfer) PalletName() string {
	return PalletName
}

func (this EventTransfer) EventIndex() uint8 {
	return 2
}

func (this EventTransfer) EventName() string {
	return "Transfer"
}

// Some amount was deposited (e.g. for transaction fees).
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

// Some amount was withdrawn from the account (e.g. for transaction fees).
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
