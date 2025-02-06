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

func (accountT) NewKeyPair(uri string) (subkey.KeyPair, error) {
	return subkey.DeriveKeyPair(sr25519.Scheme{}, uri)
}

func (accountT) GenerateAccount() (subkey.KeyPair, error) {
	return sr25519.Scheme{}.Generate()
}

func (accountT) Alice() subkey.KeyPair {
	val, err := Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice")
	if err != nil {
		panic("Should never happen.")
	}

	return val
}

func (accountT) Bob() subkey.KeyPair {
	val, err := Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Bob")
	if err != nil {
		panic("Should never happen.")
	}

	return val
}

func (accountT) Charlie() subkey.KeyPair {
	val, err := Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Charlie")
	if err != nil {
		panic("Should never happen.")
	}

	return val
}

func (accountT) Eve() subkey.KeyPair {
	val, err := Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Eve")
	if err != nil {
		panic("Should never happen.")
	}

	return val
}

func (accountT) Ferdie() subkey.KeyPair {
	val, err := Account.NewKeyPair("bottom drive obey lake curtain smoke basket hold race lonely fit walk//Ferdie")
	if err != nil {
		panic("Should never happen.")
	}

	return val
}

func (accountT) Balance(client *Client, accountId primitives.AccountId) (metadata.AccountData, error) {
	storageAt, err := client.StorageAt(primitives.None[primitives.H256]())
	if err != nil {
		return metadata.AccountData{}, err
	}

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	return val.Value.AccountData, err
}

func (accountT) Nonce(client *Client, accountId primitives.AccountId) (uint32, error) {
	storageAt, err := client.StorageAt(primitives.None[primitives.H256]())
	if err != nil {
		return uint32(0), err
	}

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	return val.Value.Nonce, err
}
