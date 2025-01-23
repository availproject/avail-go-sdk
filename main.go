package main

import (
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	val, _ := primitives.NewH256FromHexString("0x513e312001b7e288baba4b6f94bca753f0a6d6d10dcfa2ff41e52f50bc936188")

	rows := []uint32{}

	rows = append(rows, 0, 1)

	_, err := sdk.Client.Rpc.Kate.QueryRows(rows, primitives.NewSome(val))
	if err != nil {
		panic(err)
	}

}
