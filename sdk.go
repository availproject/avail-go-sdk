package main

import (
	Complex "go-sdk/complex"
	DA "go-sdk/metadata/pallets/data_availability"
	Prim "go-sdk/primitives"
	"math/big"

	"github.com/itering/scale.go/utiles/uint128"
)

type SDK struct {
	Client *Complex.Client
	Tx     Transactions
}

func NewSDK(endpoint string) SDK {
	var client = Complex.NewClient(endpoint)
	client.InitMetadata(Prim.NewNone[Prim.H256]())
	return SDK{
		Client: client,
		Tx:     newTransactions(client),
	}
}

// Temp for testing
func NewSDK2(endpoint string) SDK {
	var client = Complex.NewClient(endpoint)
	return SDK{
		Client: client,
		Tx:     newTransactions(client),
	}
}

type Transactions struct {
	DataAvailability DataAvailabilityTx
}

func newTransactions(client *Complex.Client) Transactions {
	var da = DataAvailabilityTx{Client: client}
	return Transactions{
		DataAvailability: da,
	}
}

type DataAvailabilityTx struct {
	Client *Complex.Client
}

func (this *DataAvailabilityTx) SubmitData(data []byte) Complex.Transaction {
	var call = DA.CallSubmitData{
		Data: data,
	}
	return Complex.NewTransaction(this.Client, call.ToPayload())
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
