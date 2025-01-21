package examples

import (
	SDK "github.com/nmvalera/avail-go-sdk/sdk"
)

func run_account_nonce() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	nonce, err := sdk.Client.Rpc.System.AccountNextIndex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}
	println("Nonce: ", nonce)
}
