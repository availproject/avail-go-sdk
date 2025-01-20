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
		storage := stakPal.StorageActiveEra{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value.Unwrap().Index)
		println(value.Unwrap().Start.Unwrap())
	}

	{
		storage := stakPal.StorageBonded{}
		value, err := storage.FetchAll(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value[0].Key.ToHuman())
		println(value[0].Value.ToHuman())
	}

	{
		storage := stakPal.StorageBondedEras{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value[0].Tup1)
		println(value[0].Tup2)
		println(len(value))
	}

	{
		storage := stakPal.StorageCanceledSlashPayout{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value.ToHuman())
	}

	{
		storage := stakPal.StorageChillThreshold{}
		value, err := storage.Fetch(&storageAt)
		if err != nil {
			panic(err)
		}
		println(value.IsSome())
	}

	/*
		 	println(primitives.Hex.ToHex(value.Key))
			println(value.Value.AppId)
			println(value.Value.Owner.ToHuman())
	*/
}
