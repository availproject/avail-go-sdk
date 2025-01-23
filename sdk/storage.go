package sdk

import (
	"github.com/availproject/avail-go-sdk/primitives"
)

type BlockStorage struct {
	client *Client
	at     primitives.H256
}

func (this *BlockStorage) Fetch(storageEntryKey string) (string, error) {
	return this.client.Rpc.State.GetStorage(storageEntryKey, primitives.NewSome(this.at))
}

func (this *BlockStorage) FetchKeys(storageEntryKey string) ([]string, error) {
	return this.client.Rpc.State.GetKeys(storageEntryKey, primitives.NewSome(this.at))
}
