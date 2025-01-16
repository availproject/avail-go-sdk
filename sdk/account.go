package sdk

import (
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
