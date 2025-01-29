package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	syPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountBalance() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	// Via Storage RPC
	storageAt, err := sdk.Client.StorageAt(primitives.NewNone[primitives.H256]())
	PanicOnError(err)

	storage := syPallet.StorageAccount{}
	val, err := storage.Fetch(&storageAt, accountId)
	PanicOnError(err)

	fmt.Println("Free Balance: ", val.Value.AccountData.Free.ToHuman())
	fmt.Println("Reserved Balance: ", val.Value.AccountData.Reserved.ToHuman())
	fmt.Println("Frozen Balance: ", val.Value.AccountData.Frozen.ToHuman())

	// Via Abstraction
	balance, err := SDK.Account.Balance(sdk.Client, accountId)
	PanicOnError(err)

	fmt.Println("Free Balance: ", balance.Free.ToHuman())
	fmt.Println("Reserved Balance: ", balance.Reserved.ToHuman())
	fmt.Println("Frozen Balance: ", balance.Frozen.ToHuman())

	fmt.Println("RunAccountBalance finished correctly.")
}
