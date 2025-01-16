package main

import (
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
	"go-sdk/complex"
	Balances "go-sdk/metadata/pallets/balances"
)

func main() {
	sdk := NewSDK(LocalEndpoint)
	uri := "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice"
	scheme := new(sr25519.Scheme)
	account, _ := subkey.DeriveKeyPair(scheme, uri)

	var tx = sdk.Tx.DataAvailability.SubmitData([]byte("aabbcc"))
	var options = complex.NewTransactionOptions().WithAppId(uint32(1))

	var details = tx.ExecuteAndWatch(account, false, options)
	println("Block Hash: ", details.BlockHash.ToHexWith0x())
	println("Block Number: ", details.BlockNumber)
	println("TX Events ", len(details.Events.Unwrap()))

	var events = details.Events.UnwrapOrDefault()

	var event = complex.FindFirst(events, Balances.EventWithdraw{}).UnwrapOrDefault()
	println("Who:", event.Who.ToHuman())
	println("Amount: ", event.Amount.ToHuman())

}
