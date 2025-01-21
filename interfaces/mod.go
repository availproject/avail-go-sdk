package interfaces

import (
	primitives "github.com/nmvalera/avail-go-sdk/primitives"
)

type CallDataT interface {
	Decode(call primitives.Call) primitives.Option[interface{}]
}

type EventT interface {
	PalletIndex() uint8
	PalletName() string
	EventIndex() uint8
	EventName() string
}
