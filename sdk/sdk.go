package sdk

import (
	"github.com/itering/scale.go/utiles/uint128"

	"math/big"

	"go-sdk/metadata"
	baPallet "go-sdk/metadata/pallets/balances"
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
	client           *Client
	DataAvailability DataAvailabilityTx
	Utility          UtilityTx
	Balances         BalancesTx
}

func newTransactions(client *Client) Transactions {
	return Transactions{
		client:           client,
		DataAvailability: DataAvailabilityTx{client: client},
		Utility:          UtilityTx{client: client},
		Balances:         BalancesTx{client: client},
	}
}

type DataAvailabilityTx struct {
	client *Client
}

func (this *Transactions) NewTransaction(payload metadata.Payload) Transaction {
	return NewTransaction(this.client, payload)
}

func (this *DataAvailabilityTx) SubmitData(data []byte) Transaction {
	call := daPallet.CallSubmitData{Data: data}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *DataAvailabilityTx) CreateApplicationKey(key []byte) Transaction {
	call := daPallet.CallCreateApplicationKey{Key: key}
	return NewTransaction(this.client, call.ToPayload())
}

type UtilityTx struct {
	client *Client
}

func (this *UtilityTx) Batch(calls []prim.Call) Transaction {
	call := utPallet.CallBatch{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *UtilityTx) BatchAll(calls []prim.Call) Transaction {
	call := utPallet.CallBatchAll{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *UtilityTx) ForceBatch(calls []prim.Call) Transaction {
	call := utPallet.CallForceBatch{Calls: calls}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *UtilityTx) AsDerivate(index uint16, call prim.Call) Transaction {
	c := utPallet.CallAsDerivate{Index: index, Call: call}
	return NewTransaction(this.client, c.ToPayload())
}

type BalancesTx struct {
	client *Client
}

func (this *BalancesTx) TransferAllowDeath(dest metadata.AccountId, amount uint128.Uint128) Transaction {
	call := baPallet.CallTransferAlowDeath{Dest: dest.ToMultiAddress(), Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *BalancesTx) ForceTransfer(dest metadata.AccountId, amount uint128.Uint128) Transaction {
	call := baPallet.CallForceTransfer{Dest: dest.ToMultiAddress(), Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

func (this *BalancesTx) TransferKeepAlive(dest metadata.AccountId, amount uint128.Uint128) Transaction {
	call := baPallet.CallTransferKeepAlive{Dest: dest.ToMultiAddress(), Value: amount}
	return NewTransaction(this.client, call.ToPayload())
}

func OneAvail() uint128.Uint128 {
	var res, _ = new(big.Int).SetString("1000000000000000000", 10)
	return uint128.FromBig(res)
}

const LocalEndpoint = "http://127.0.0.1:9944"
const TuringEndpoint = "https://turing-rpc.avail.so/rpc"
const MainnetEndpoint = "https://mainnet-rpc.avail.so/rpc"
