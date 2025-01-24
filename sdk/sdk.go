package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type SDK struct {
	Client *Client
	Tx     Transactions
}

func (this *SDK) UpdateMetadata(blockHash prim.Option[prim.H256]) error {
	return this.Client.InitMetadata(blockHash)
}

// Returns a new SDK using the latest metadata from the chain.
// To get the SDK initialized with different metadata, call NewSDKWithMetadata#
//
// In 99% cases this is the one that you need to call. In case you are exploring
// historical blocks that needs different metadata then make sure to call
// NewSDKWithMetadata instead of this.
//
// The metadata can be updated on fly by calling sdk.UpdateMetadata(blockHash)
func NewSDK(endpoint string) (SDK, error) {
	return NewSDKWithMetadata(endpoint, prim.NewNone[prim.H256]())
}

// Same as NewSDK but allows passing the block hash from which the metadata will be
// fetched.
func NewSDKWithMetadata(endpoint string, metadataBlockHash prim.Option[prim.H256]) (SDK, error) {
	var client = NewClient(endpoint)

	// Temp for testing
	if err := client.InitMetadata(metadataBlockHash); err != nil {
		return SDK{}, nil
	}
	return SDK{
		Client: client,
		Tx:     newTransactions(client),
	}, nil
}

type Transactions struct {
	client           *Client
	DataAvailability DataAvailabilityTx
	Utility          UtilityTx
	Balances         BalancesTx
	Staking          StakingTx
	NominationPools  NominationPoolsTx
	System           SystemTx
	Vector           VectorTx
}

func newTransactions(client *Client) Transactions {
	return Transactions{
		client:           client,
		DataAvailability: DataAvailabilityTx{client: client},
		Utility:          UtilityTx{client: client},
		Balances:         BalancesTx{client: client},
		Staking:          StakingTx{client: client},
		NominationPools:  NominationPoolsTx{client: client},
		System:           SystemTx{client: client},
		Vector:           VectorTx{client: client},
	}
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
