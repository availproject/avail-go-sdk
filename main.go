package main

import (
	// "go-sdk/examples"

	npmPool "go-sdk/metadata/pallets/nomination_pools"
	"go-sdk/primitives"
	SDK "go-sdk/sdk"
	/*
		 	"go-sdk/metadata"
			"go-sdk/primitives"

	*/)

func main() {
	sdk := SDK.NewClient(SDK.TuringEndpoint)

	storageAt, err := sdk.StorageAt(primitives.NewNone[primitives.H256]())
	if err != nil {
		panic(err)
	}

	{
		storage := npmPool.StorageTotalValueLocked{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value.String())
	}
}
