package main

import (
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
	"go-sdk/complex"

	balancesPallet "go-sdk/metadata/pallets/balances"
)

func main() {
	sdk := complex.NewSDK(complex.LocalEndpoint)
	uri := "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice"
	scheme := new(sr25519.Scheme)
	account, _ := subkey.DeriveKeyPair(scheme, uri)

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("aabbcc"))
	options := complex.NewTransactionOptions().WithAppId(uint32(1))

	details, err := tx.ExecuteAndWatchInclusion(account, options)
	if err != nil {
		panic(err)
	}
	println("Block Hash: ", details.BlockHash.ToHexWith0x())
	println("Block Number: ", details.BlockNumber)
	println("TX Events ", len(details.Events.Unwrap()))

	var events = details.Events.Unwrap()

	var event = complex.FindFirst(events, balancesPallet.EventWithdraw{}).Unwrap()
	println("Who:", event.Who.ToHuman())
	println("Amount: ", event.Amount.ToHuman())

}
