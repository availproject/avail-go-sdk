package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	prim "go-sdk/primitives"
)

type SDK struct {
	Client *Client
	Tx     Transactions
}

func NewSDK(endpoint string) SDK {
	var client = NewClient(endpoint)

	// Temp for testing
	if err := client.InitMetadata(prim.NewNone[prim.H256]()); err != nil {
		panic(err)
	}
	return SDK{
		Client: client,
		Tx:     newTransactions(client),
	}
}

// Temp for testing
func NewSDK2(endpoint string) SDK {
	var client = NewClient(endpoint)
	return SDK{
		Client: client,
		Tx:     newTransactions(client),
	}
}

type Transactions struct {
	client           *Client
	DataAvailability DataAvailabilityTx
	Utility          UtilityTx
	Balances         BalancesTx
	Staking          StakingTx
}

func newTransactions(client *Client) Transactions {
	return Transactions{
		client:           client,
		DataAvailability: DataAvailabilityTx{client: client},
		Utility:          UtilityTx{client: client},
		Balances:         BalancesTx{client: client},
		Staking:          StakingTx{client: client},
	}
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
