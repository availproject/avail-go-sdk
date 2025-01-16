package main

import (
	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"

	balancesPallet "go-sdk/metadata/pallets/balances"
	SDK "go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)
	uri := "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice"
	account, err := subkey.DeriveKeyPair(sr25519.Scheme{}, uri)
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("ThisIsMyData"))
	options := SDK.NewTransactionOptions().WithAppId(uint32(2))

	details, err := tx.ExecuteAndWatchInclusion(account, options)
	if err != nil {
		panic(err)
	}
	println("Block Hash: ", details.BlockHash.ToHexWith0x())
	println("Block Number: ", details.BlockNumber)
	println("TX Events ", len(details.Events.Unwrap()))
	var events = details.Events.Unwrap()

	event := SDK.EventFindFirst(events, balancesPallet.EventWithdraw{}).Unwrap()
	println("Who:", event.Who.ToHuman())
	println("Amount: ", event.Amount.ToHuman())

	block, err := SDK.NewBlock(sdk.Client, details.BlockHash)
	if err != nil {
		panic(err)
	}

	res1 := block.DataSubmissionByAppId(2)
	println("Len of by app Id:", len(res1))
	println("Data: ", string(block.DataSubmissionByIndex(details.TxIndex).Unwrap().Data))

}
