package sdk

import (
	"github.com/availproject/avail-go-sdk/primitives"
)

type BlockStorage struct {
	client *Client
	at     primitives.H256
}

func (this *BlockStorage) Fetch(storageEntryKey string) (primitives.Option[string], error) {
	return this.client.Rpc.State.GetStorage(storageEntryKey, primitives.Some(this.at))
}

func (this *BlockStorage) FetchKeys(storageEntryKey string) ([]string, error) {
	return this.client.Rpc.State.GetKeys(storageEntryKey, primitives.Some(this.at))
}

func (this *BlockStorage) FetchKeysPaged(storageEntryKey string, count uint32, startKey primitives.Option[string]) ([]string, error) {
	return this.client.Rpc.State.GetKeysPaged(storageEntryKey, count, startKey, primitives.Some(this.at))
}
