package interfaces

import "github.com/availproject/avail-go-sdk/primitives"

type EventT interface {
	PalletIndex() uint8
	PalletName() string
	EventIndex() uint8
	EventName() string
}

type BlockStorageT interface {
	Fetch(key string) (primitives.Option[string], error)
	FetchKeys(key string) ([]string, error)
	FetchKeysPaged(key string, count uint32, startKey primitives.Option[string]) ([]string, error)
}
