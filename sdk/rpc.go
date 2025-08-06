package sdk

import (
	"fmt"
	"strconv"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type RPC struct {
	client      *Client
	System      systemRPC
	State       stateRPC
	Chain       chainRPC
	ChainSpec   chainSpecRPC
	Kate        kateRPC
	Author      authorRPC
	Transaction transactionRPC
}

func newRPC(client *Client) RPC {
	return RPC{
		client:      client,
		System:      systemRPC{client: client},
		State:       stateRPC{client: client},
		Chain:       chainRPC{client: client},
		ChainSpec:   chainSpecRPC{client: client},
		Kate:        kateRPC{client: client},
		Author:      authorRPC{client: client},
		Transaction: transactionRPC{client: client},
	}
}

type RPCParams struct {
	Values []string
}

func (r *RPCParams) Add(value string) {
	r.Values = append(r.Values, value)
}

func (r *RPCParams) AddBool(value bool) {
	if value == true {
		r.Add("true")
	} else {
		r.Add("false")
	}
}

func (r *RPCParams) AddByteSlice(value []byte) {
	if len(value) == 0 {
		return
	}

	res := "["
	for i := range value {
		res += fmt.Sprintf("%v", value[i])

		if i < (len(value) - 1) {
			res += ","
		}
	}

	res = res + "]"

	r.Values = append(r.Values, res)
}

func (r *RPCParams) AddH256(value prim.H256) {
	r.Add(value.ToRpcParam())
}

func (r *RPCParams) AddUint32(value uint32) {
	r.Add(strconv.FormatUint(uint64(value), 10))
}

func (r *RPCParams) Build() string {
	length := len(r.Values)
	if length == 0 {
		return "[]"
	}

	result := "["
	for i := 0; i < length; i++ {
		result += r.Values[i]
		if i < (length - 1) {
			result += ", "
		}
	}
	result += "]"

	return result
}
