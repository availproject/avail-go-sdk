package main

import (
	"go-sdk/metadata"
	baPallet "go-sdk/metadata/pallets/balances"
	"go-sdk/primitives"
	SDK "go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

	acc2, err := metadata.NewAccountIdFromAddress("5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty")
	if err != nil {
		panic(err)
	}

	multiaddres := primitives.NewMultiAddressId(acc2.Value)
	amount := SDK.OneAvail()

	call1 := baPallet.CallTransferKeepAlive{
		Dest:  multiaddres,
		Value: amount,
	}

	calls := []primitives.Call{}

	calls = append(calls, call1.ToCall())
	calls = append(calls, call1.ToCall())
	calls = append(calls, call1.ToCall())

	tx := sdk.Tx.Utility.Batch(calls)
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(0))
	if err != nil {
		panic(err)
	}

	println(res.BlockHash.ToHuman())

	println("All Good :)")
}
