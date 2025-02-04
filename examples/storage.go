package examples

import (
	"fmt"

	idenPallet "github.com/availproject/avail-go-sdk/metadata/pallets/identity"
	staPallet "github.com/availproject/avail-go-sdk/metadata/pallets/staking"
	sysPallet "github.com/availproject/avail-go-sdk/metadata/pallets/system"
	prim "github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunStorage() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := prim.NewH256FromHexString("0x9e813bb85fca217f8f3967bd4b550b05f7d559412571ca1dd621aa37343b300b")
	PanicOnError(err)

	blockStorage, err := sdk.Client.StorageAt(prim.NewSome(blockHash))
	PanicOnError(err)

	// Simple Storage
	{
		storage := staPallet.StorageMinValidatorBond{}
		val, err := storage.Fetch(&blockStorage)
		PanicOnError(err)

		fmt.Println("Min Validator Bond: ", val.ToHuman())
	}

	// Simple Storage that returns Option
	{
		storage := staPallet.StorageCurrentEra{}
		val, err := storage.Fetch(&blockStorage)
		PanicOnError(err)

		if val.IsSome() {
			fmt.Println("Current Era: ", val.Unwrap())
		}
	}

	// Fetch Map Storage
	{
		storage := sysPallet.StorageAccount{}
		acc, err := prim.NewAccountIdFromAddress("5C869t2dWzmmYkE8NT1oocuEEdwqNnAm2XhvnuHcavNUcTTT")
		PanicOnError(err)

		val, err := storage.Fetch(&blockStorage, acc)
		PanicOnError(err)

		fmt.Println("Account Key: ", val.Key.ToHuman())
		fmt.Println("Account Nonce: ", val.Value.Nonce)
		fmt.Println("Account Free Balance: ", val.Value.AccountData.Free.ToHuman())
	}

	// Fetch All Map Storage
	{
		storage := idenPallet.StorageIdentityOf{}
		val, err := storage.FetchAll(&blockStorage)
		PanicOnError(err)
		AssertTrue(len(val) > 0, "There need to be more than 0 values")

		for i := 0; i < len(val); i++ {
			fmt.Println("Identity Key: ", val[i].Key.ToHuman())
			fmt.Println("Identity Deposit: ", val[i].Value.T0.Deposit.ToHuman())
			fmt.Println("Identity Display: ", val[i].Value.T0.Info.Display.ToHuman())
			if i >= 2 {
				break
			}
		}
	}

	// Fetch Double Map Storage
	{
		storage := staPallet.StorageErasValidatorPrefs{}
		era := uint32(299)
		acc, err := prim.NewAccountIdFromAddress("5EFTSpRN2nMZDLjkniBYdmMxquMNm5CLVsrX2V3HHue6QFFF")
		PanicOnError(err)

		val, err := storage.Fetch(&blockStorage, era, acc)
		PanicOnError(err)

		fmt.Println("Era: ", val.Key1)
		fmt.Println("Address: ", val.Key2.ToHuman())
		fmt.Println("Commission: ", val.Value.Commission.ToHuman())
		fmt.Println("Blocked: ", val.Value.Blocked)
	}

	// Fetch All Double Map Storage
	{
		storage := staPallet.StorageErasValidatorPrefs{}
		era := uint32(299)
		val, err := storage.FetchAll(&blockStorage, era)
		PanicOnError(err)
		AssertTrue(len(val) > 0, "There need to be more than 0 values")

		for i := 0; i < len(val); i++ {
			fmt.Println("Era: ", val[i].Key1)
			fmt.Println("Address: ", val[i].Key2.ToHuman())
			fmt.Println("Commission: ", val[i].Value.Commission.ToHuman())
			fmt.Println("Blocked: ", val[i].Value.Blocked)
			if i >= 2 {
				break
			}
		}
	}

	fmt.Println("RunStorage finished correctly.")
}
