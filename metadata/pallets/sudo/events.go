package sudo

import (
	"github.com/availproject/avail-go-sdk/metadata"
)

// A sudo call just took place.
type EventSudid struct {
	SudoResult metadata.DispatchResult
}

func (this EventSudid) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventSudid) PalletName() string {
	return PalletName
}

func (this EventSudid) EventIndex() uint8 {
	return 0
}

func (this EventSudid) EventName() string {
	return "Sudid"
}

// A [sudo_as](Pallet::sudo_as) call just took place.
type EventSudoAsDone struct {
	SudoResult metadata.DispatchResult
}

func (this EventSudoAsDone) PalletIndex() uint8 {
	return PalletIndex
}

func (this EventSudoAsDone) PalletName() string {
	return PalletName
}

func (this EventSudoAsDone) EventIndex() uint8 {
	return 3
}

func (this EventSudoAsDone) EventName() string {
	return "SudoAsDone"
}
