package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	stPallet "github.com/availproject/avail-go-sdk/metadata/pallets/staking"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunValidator() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	// Generating new account.
	acc, err := SDK.Account.GenerateAccount()
	PanicOnError(err)

	// Sending funds to that account.
	dest := metadata.NewAccountIdFromKeyPair(acc).ToMultiAddress()
	tx := sdk.Tx.Balances.TransferKeepAlive(dest, SDK.OneAvail().Mul64(uint64(250_000)))
	res, err := tx.ExecuteAndWatchInclusion(SDK.Account.Alice(), SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction must be successful")

	// Fetching Min Validator Bond storage
	blockStorage, err := sdk.Client.StorageAt(primitives.NewNone[primitives.H256]())
	PanicOnError(err)

	storage := stPallet.StorageMinValidatorBond{}
	minValBond, err := storage.Fetch(&blockStorage)
	PanicOnError(err)

	// If there is a min validator bond value then we will bond 1 more.
	// If there isn't one then instead of bonding 0 we will bond 1.
	bondValue := minValBond.Add(SDK.OneAvail())
	payee := metadata.RewardDestination{VariantIndex: 0}

	// Bond
	tx = sdk.Tx.Staking.Bond(bondValue, payee)
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction must be successful")

	// Generate Session Keys
	keysRaw, err := sdk.Client.Rpc.Author.RotateKeys()
	PanicOnError(err)
	sessionKeys, err := SDK.DeconstructSessionKeys(keysRaw)
	PanicOnError(err)

	// Set Keys
	tx = sdk.Tx.SessionTx.SetKeys(sessionKeys, []byte{})
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction must be successful")

	// Validate
	commission := metadata.NewPerbillFromU8(10) // 10.0%
	pref := metadata.ValidatorPrefs{Commission: commission, Blocked: false}

	tx = sdk.Tx.Staking.Validate(pref)
	res, err = tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction must be successful")

	fmt.Println("RunValidator finished correctly.")

}
