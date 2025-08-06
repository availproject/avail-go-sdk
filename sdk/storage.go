package sdk

import (
	"github.com/availproject/avail-go-sdk/primitives"
)

type BlockStorage struct {
	client *Client
	at     primitives.H256
}

func (b *BlockStorage) Fetch(storageEntryKey string) (primitives.Option[string], error) {
	return b.client.Rpc.State.GetStorage(storageEntryKey, primitives.Some(b.at))
}

func (b *BlockStorage) FetchKeys(storageEntryKey string) ([]string, error) {
	return b.client.Rpc.State.GetKeys(storageEntryKey, primitives.Some(b.at))
}

func (b *BlockStorage) FetchKeysPaged(storageEntryKey string, count uint32, startKey primitives.Option[string]) ([]string, error) {
	return b.client.Rpc.State.GetKeysPaged(storageEntryKey, count, startKey, primitives.Some(b.at))
}
