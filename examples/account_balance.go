package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func Run_account_balance() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	// Via Abstraction
	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}
	balance, err := SDK.Account.Balance(sdk.Client, accountId)
	if err != nil {
		panic(err)
	}
	println("Free Balance: ", balance.Free.ToHuman())
	println("Reserved Balance: ", balance.Reserved.ToHuman())
	println("Frozen Balance: ", balance.Frozen.ToHuman())
}
