package main

import (
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.LocalEndpoint)

	val, _ := primitives.NewH256FromHexString("0xea552116539b130effd3404e409a3a9f99c55e47bcb855320bb70bc640b4eab3")

	cells := []SDK.KateCell{}

	cells = append(cells, SDK.KateCell{Row: 0, Col: 0})
	cells = append(cells, SDK.KateCell{Row: 0, Col: 1})

	_, err := sdk.Client.Rpc.Kate.QueryProof(cells, primitives.NewSome(val))
	if err != nil {
		panic(err)
	}

}
