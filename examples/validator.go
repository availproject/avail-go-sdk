package examples

import (
	"github.com/availproject/avail-go-sdk/metadata"
	stPallet "github.com/availproject/avail-go-sdk/metadata/pallets/staking"
	"github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunValidator() {
	sdk, err := SDK.NewSDK(SDK.LocalEndpoint)
	PanicOnError(err)

	acc := SDK.Account.Bob()

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
	tx := sdk.Tx.Staking.Bond(bondValue, payee)
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions())
	PanicOnError(err)
	AssertTrue(res.IsSuccessful().Unwrap(), "Transaction must be successful")

	// Generate Session Keys
	keysRaw, err := sdk.Client.Rpc.Author.RotateKeys()
	PanicOnError(err)
	_, err = SDK.DeconstructSessionKeys(keysRaw)
	PanicOnError(err)

	// Set Keys

}
