package main

import (
	// "go-sdk/examples"

	stakPal "go-sdk/metadata/pallets/staking"
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

	{
		storage := stakPal.StorageMinimumActiveStake{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value.ToHuman())
	}

	{
		storage := stakPal.StorageMinimumValidatorCount{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value)
	}

	/*
		 	println(primitives.Hex.ToHex(value.Key))
			println(value.Value.AppId)
			println(value.Value.Owner.ToHuman())
	*/
}
