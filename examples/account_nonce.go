package examples

import (
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func Run_account_nonce() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	nonce, err := sdk.Client.Rpc.System.AccountNextIndex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}
	println("Nonce: ", nonce)
}
