package main

import (
	//"github.com/availproject/avail-go-sdk/examples"

	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	if err != nil {
		panic(err)
	}

	acc := SDK.Account.Alice()

	accountId := metadata.NewAccountIdFromKeyPair(acc)
	val, err := SDK.Account.Balance(sdk.Client, accountId)
	if err != nil {
		panic(err)
	}
	println(val.Free.ToHuman())

	val2, err := SDK.Account.Nonce(sdk.Client, accountId)
	if err != nil {
		panic(err)
	}
	println(val2)

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.TransactionOptions{}.WithAppId(1))
	if err != nil {
		panic(err)
	}
	if isOk, err := res.IsSuccessful(); err != nil {
		panic(err)
	} else if !isOk {
		panic("Tx Failed.")
	}
}
