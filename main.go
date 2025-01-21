package main

import (
	// "go-sdk/examples"

	"go-sdk/metadata"
	"go-sdk/primitives"
	SDK "go-sdk/sdk"
	/*
		 	"go-sdk/metadata"
			"go-sdk/primitives"

	*/)

func main() {
	sdk := SDK.NewClient(SDK.LocalEndpoint)

	storageAt, err := sdk.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		panic(err)
	}

	println(&storageAt)

	val := metadata.Perbill{}
	val.Value = 10_00_000
	println(val.ToHuman())

	/*
		 	println(primitives.Hex.ToHex(value.Key))
			println(value.Value.AppId)
			println(value.Value.Owner.ToHuman())
	*/
}
