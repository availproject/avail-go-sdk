package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	"go-sdk/metadata"
	daPallet "go-sdk/metadata/pallets/data_availability"
	utPallet "go-sdk/metadata/pallets/utility"
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
	Client           *Client
	DataAvailability DataAvailabilityTx
	Utility          UtilityTx
}

func newTransactions(client *Client) Transactions {
	return Transactions{
		Client:           client,
		DataAvailability: DataAvailabilityTx{Client: client},
		Utility:          UtilityTx{Client: client},
	}
}

type DataAvailabilityTx struct {
	Client *Client
}

func (this *Transactions) NewTransaction(payload metadata.Payload) Transaction {
	return NewTransaction(this.Client, payload)
}

func (this *DataAvailabilityTx) SubmitData(data []byte) Transaction {
	call := daPallet.CallSubmitData{Data: data}
	return NewTransaction(this.Client, call.ToPayload())
}

func (this *DataAvailabilityTx) CreateApplicationKey(key []byte) Transaction {
	call := daPallet.CallCreateApplicationKey{Key: key}
	return NewTransaction(this.Client, call.ToPayload())
}

type UtilityTx struct {
	Client *Client
}

func (this *UtilityTx) Batch(calls []prim.Call) Transaction {
	call := utPallet.CallBatch{Calls: calls}
	return NewTransaction(this.Client, call.ToPayload())
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
