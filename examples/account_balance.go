package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountBalance() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}

	// Via Storage RPC
	storageAt, err := sdk.Client.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		panic(err)
	}

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	if err != nil {
		panic(err)
	}

	println("Free Balance: ", val.Value.AccountData.Free.ToHuman())
	println("Reserved Balance: ", val.Value.AccountData.Reserved.ToHuman())
	println("Frozen Balance: ", val.Value.AccountData.Frozen.ToHuman())

	// Via Abstraction
	balance, err := SDK.Account.Balance(sdk.Client, accountId)
	if err != nil {
		panic(err)
	}
	println("Free Balance: ", balance.Free.ToHuman())
	println("Reserved Balance: ", balance.Reserved.ToHuman())
	println("Frozen Balance: ", balance.Frozen.ToHuman())

	println("RunAccountBalance finished correctly.")
}
