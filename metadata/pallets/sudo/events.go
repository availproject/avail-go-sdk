package sudo

import (
	"github.com/availproject/avail-go-sdk/metadata"
)

// A sudo call just took place.
type EventSudid struct {
	SudoResult metadata.DispatchResult
}

func (es EventSudid) PalletIndex() uint8 {
	return PalletIndex
}

func (es EventSudid) PalletName() string {
	return PalletName
}

func (es EventSudid) EventIndex() uint8 {
	return 0
}

func (es EventSudid) EventName() string {
	return "Sudid"
}

// A [sudo_as](Pallet::sudo_as) call just took place.
type EventSudoAsDone struct {
	SudoResult metadata.DispatchResult
}

func (esad EventSudoAsDone) PalletIndex() uint8 {
	return PalletIndex
}

func (esad EventSudoAsDone) PalletName() string {
	return PalletName
}

func (esad EventSudoAsDone) EventIndex() uint8 {
	return 3
}

func (esad EventSudoAsDone) EventName() string {
	return "SudoAsDone"
}
