package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountNonce() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	// Via RPC
	nonce, err := sdk.Client.Rpc.System.AccountNextIndex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}
	println("RPC Nonce: ", nonce)

	// Via Abstraction
	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	if err != nil {
		panic(err)
	}
	nonce2, err := SDK.Account.Nonce(sdk.Client, accountId)
	if err != nil {
		panic(err)
	}
	println("Abstraction Nonce: ", nonce2)
}
