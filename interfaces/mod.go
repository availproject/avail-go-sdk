package interfaces

import ()

type EventT interface {
	PalletIndex() uint8
	PalletName() string
	EventIndex() uint8
	EventName() string
}

type BlockStorageT interface {
	Fetch(storageEntryKey string) (string, error)
	FetchKeys(key string) ([]string, error)
}
