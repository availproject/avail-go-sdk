package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountNonce() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	// Via RPC
	nonce, err := sdk.Client.Rpc.System.AccountNextIndex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	println("RPC Nonce: ", nonce)

	// Via Abstraction
	accountId, err := metadata.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)

	nonce2, err := SDK.Account.Nonce(sdk.Client, accountId)
	PanicOnError(err)

	println("Abstraction Nonce: ", nonce2)

	println("RunAccountNonce finished correctly.")
}
