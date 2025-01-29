package sdk

import (
	"github.com/availproject/avail-go-sdk/metadata"
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type RuntimeApi struct {
	client *Client
}

func newRuntimeAPi(client *Client) RuntimeApi {
	return RuntimeApi{
		client: client,
	}
}

// Parameter "tx" is hex encoded transaction.
func (this *RuntimeApi) TransactionPaymentApi_queryInfo(tx string, at prim.Option[prim.H256]) (metadata.RuntimeDispatchInfo, error) {
	encodedTxLen := len(prim.Hex.FromHex(tx))
	encoded := tx + prim.Encoder.Encode(uint32(encodedTxLen))

	val, err := this.client.Rpc.State.Call("TransactionPaymentApi_query_info", encoded, at)
	if err != nil {
		return metadata.RuntimeDispatchInfo{}, err
	}

	res := metadata.RuntimeDispatchInfo{}
	decoder := prim.NewDecoder(prim.Hex.FromHex(val), 0)
	err = decoder.Decode(&res)
	return res, err
}

// Parameter "tx" is hex encoded transaction.
func (this *RuntimeApi) TransactionPaymentApi_queryFeeDetails(tx string, at prim.Option[prim.H256]) (metadata.FeeDetails, error) {
	encodedTxLen := len(prim.Hex.FromHex(tx))
	encoded := tx + prim.Encoder.Encode(uint32(encodedTxLen))

	val, err := this.client.Rpc.State.Call("TransactionPaymentApi_query_fee_details", encoded, at)
	if err != nil {
		return metadata.FeeDetails{}, err
	}

	res := metadata.FeeDetails{}
	decoder := prim.NewDecoder(prim.Hex.FromHex(val), 0)
	err = decoder.Decode(&res)
	return res, err
}

// Parameter "call" is hex encoded "primitive.call".
func (this *RuntimeApi) TransactionPaymentCallApi_queryCallInfo(call string, at prim.Option[prim.H256]) (metadata.RuntimeDispatchInfo, error) {
	encodedTxLen := len(prim.Hex.FromHex(call))
	encoded := call + prim.Encoder.Encode(uint32(encodedTxLen))

	val, err := this.client.Rpc.State.Call("TransactionPaymentCallApi_query_call_info", encoded, at)
	if err != nil {
		return metadata.RuntimeDispatchInfo{}, err
	}

	res := metadata.RuntimeDispatchInfo{}
	decoder := prim.NewDecoder(prim.Hex.FromHex(val), 0)
	err = decoder.Decode(&res)
	return res, err
}

// Parameter "call" is hex encoded "primitive.call".
func (this *RuntimeApi) TransactionPaymentCallApi_queryCallFeeDetails(call string, at prim.Option[prim.H256]) (metadata.FeeDetails, error) {
	encodedTxLen := len(prim.Hex.FromHex(call))
	encoded := call + prim.Encoder.Encode(uint32(encodedTxLen))

	val, err := this.client.Rpc.State.Call("TransactionPaymentCallApi_query_call_fee_details", encoded, at)
	if err != nil {
		return metadata.FeeDetails{}, err
	}

	res := metadata.FeeDetails{}
	decoder := prim.NewDecoder(prim.Hex.FromHex(val), 0)
	err = decoder.Decode(&res)
	return res, err
}
