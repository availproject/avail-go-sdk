package sdk

import (
	"github.com/availproject/avail-go-sdk/metadata"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	"github.com/availproject/avail-go-sdk/primitives"

	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
)

type accountT struct{}

var Account accountT

func (accountT) NewKeyPair(uri string) (kp subkey.KeyPair, err error) {
	return subkey.DeriveKeyPair(sr25519.Scheme{}, uri)
}

func (accountT) Alice() (kp subkey.KeyPair, err error) {
	return Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice")
}

func (accountT) Balance(client *Client, accountId metadata.AccountId) (metadata.AccountData, error) {
	storageAt, err := client.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		return metadata.AccountData{}, err
	}

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	return val.Value.AccountData, err
}

func (accountT) Nonce(client *Client, accountId metadata.AccountId) (uint32, error) {
	storageAt, err := client.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		return uint32(0), err
	}

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	return val.Value.Nonce, err
}
