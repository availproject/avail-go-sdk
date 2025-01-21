package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	daPallet "github.com/availproject/avail-go-sdk/metadata/pallets/data_availability"
	prim "github.com/availproject/avail-go-sdk/primitives"
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
	DataAvailability DataAvailabilityTx
}

func newTransactions(client *Client) Transactions {
	return Transactions{
		DataAvailability: DataAvailabilityTx{Client: client},
	}
}

type DataAvailabilityTx struct {
	Client *Client
}

func (this *DataAvailabilityTx) SubmitData(data []byte) Transaction {
	call := daPallet.CallSubmitData{Data: data}
	return NewTransaction(this.Client, call.ToPayload())
}

func (this *DataAvailabilityTx) CreateApplicationKey(key []byte) Transaction {
	call := daPallet.CallCreateApplicationKey{Key: key}
	return NewTransaction(this.Client, call.ToPayload())
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
