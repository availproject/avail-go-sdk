package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunAccountCreation() {
	// Use `NewKeyPair` to create your own key from your secret seed phase or uri.
	acc, err := SDK.Account.NewKeyPair("//Alice")
	PanicOnError(err)
	fmt.Println("Alice Address: " + acc.SS58Address(42))

	// Use `GenerateAccount` to generate a random account
	acc, err = SDK.Account.GenerateAccount()
	PanicOnError(err)
	fmt.Println("Random Account Address: " + acc.SS58Address(42))

	// There are predefined testing accounts available to be used on local dev networks.
	acc = SDK.Account.Alice()
	fmt.Println("Alice Address: " + acc.SS58Address(42))
	acc = SDK.Account.Bob()
	fmt.Println("Bob Address: " + acc.SS58Address(42))
	acc = SDK.Account.Charlie()
	fmt.Println("Charlie Address: " + acc.SS58Address(42))
	acc = SDK.Account.Eve()
	fmt.Println("Eve Address: " + acc.SS58Address(42))
	acc = SDK.Account.Ferdie()
	fmt.Println("Ferdie Address: " + acc.SS58Address(42))

	// AccountId can be created from Keypair...
	accountId := primitives.NewAccountIdFromKeyPair(acc)
	fmt.Println("Ferdie Address: " + accountId.ToHuman())

	// ...or from SS58 address
	accountId, err = primitives.NewAccountIdFromAddress("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	PanicOnError(err)
	fmt.Println("Alice Address: " + accountId.ToHuman())
	fmt.Println("Alice Account Id: " + accountId.ToString())

	// AccountId can be converted to MultiAddress
	multiAddress := accountId.ToMultiAddress()
	AssertEq(multiAddress.VariantIndex, 0, "Variant Index needs to be 0")
	AssertTrue(multiAddress.Id.IsSome(), "ID needs to be populated")
	AssertEq(multiAddress.Id.UnsafeUnwrap(), accountId, "multiAddress and accountId need to have the same value")

	// MultiAddress can be converted to AccountId
	accountId2 := multiAddress.ToAccountId()
	AssertTrue(accountId2.IsSome(), "")

	// Non-init AccountId has `5C4hrfjw9DjXZTzV3MwzrrAr9P1MJhSrvWGWqi1eSuyUpnhM` as the SS58 address
	accountId = primitives.AccountId{Value: primitives.H256{}}
	fmt.Println("Address: " + accountId.ToHuman())
	AssertEq(accountId.ToHuman(), "5C4hrfjw9DjXZTzV3MwzrrAr9P1MJhSrvWGWqi1eSuyUpnhM", "Non-init account id ss58 address is not correct.")

	fmt.Println("RunAccountCreation finished correctly.")
}
