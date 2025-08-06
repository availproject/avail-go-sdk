package balances

import (
	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"
)

// An account was created with some free balance.
type EventEndowed struct {
	Account     primitives.AccountId
	FreeBalance metadata.Balance
}

func (ee EventEndowed) PalletIndex() uint8 {
	return PalletIndex
}

func (ee EventEndowed) PalletName() string {
	return PalletName
}

func (ee EventEndowed) EventIndex() uint8 {
	return 0
}

func (ee EventEndowed) EventName() string {
	return "Endowed"
}

// An account was removed whose balance was non-zero but below ExistentialDeposit, resulting in an outright loss.
type EventDustLost struct {
	Account primitives.AccountId
	Amount  metadata.Balance
}

func (edl EventDustLost) PalletIndex() uint8 {
	return PalletIndex
}

func (edl EventDustLost) PalletName() string {
	return PalletName
}

func (edl EventDustLost) EventIndex() uint8 {
	return 1
}

func (edl EventDustLost) EventName() string {
	return "DustLost"
}

// Transfer succeeded.
type EventTransfer struct {
	From   primitives.AccountId
	To     primitives.AccountId
	Amount metadata.Balance
}

func (et EventTransfer) PalletIndex() uint8 {
	return PalletIndex
}

func (et EventTransfer) PalletName() string {
	return PalletName
}

func (et EventTransfer) EventIndex() uint8 {
	return 2
}

func (et EventTransfer) EventName() string {
	return "Transfer"
}

// Some amount was deposited (e.g. for transaction fees).
type EventDeposit struct {
	Who    primitives.AccountId
	Amount metadata.Balance
}

func (ed EventDeposit) PalletIndex() uint8 {
	return PalletIndex
}

func (ed EventDeposit) PalletName() string {
	return PalletName
}

func (ed EventDeposit) EventIndex() uint8 {
	return 7
}

func (ed EventDeposit) EventName() string {
	return "Deposit"
}

// Some amount was withdrawn from the account (e.g. for transaction fees).
type EventWithdraw struct {
	Who    primitives.AccountId
	Amount metadata.Balance
}

func (ew EventWithdraw) PalletIndex() uint8 {
	return PalletIndex
}

func (ew EventWithdraw) PalletName() string {
	return PalletName
}

func (ew EventWithdraw) EventIndex() uint8 {
	return 8
}

func (ew EventWithdraw) EventName() string {
	return "Withdraw"
}
