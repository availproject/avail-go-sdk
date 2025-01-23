package sdk

import (
	"strconv"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type RPC struct {
	client    *Client
	System    systemRPC
	State     stateRPC
	Chain     chainRPC
	ChainSpec chainSpecRPC
	Kate      kateRPC
}

func newRPC(client *Client) RPC {
	return RPC{
		client:    client,
		System:    systemRPC{client: client},
		State:     stateRPC{client: client},
		Chain:     chainRPC{client: client},
		ChainSpec: chainSpecRPC{client: client},
		Kate:      kateRPC{client: client},
	}
}

type RPCParams struct {
	Values []string
}

func (this *RPCParams) Add(value string) {
	this.Values = append(this.Values, value)
}

func (this *RPCParams) AddH256(value prim.H256) {
	this.Add(value.ToRpcParam())
}

func (this *RPCParams) AddUint32(value uint32) {
	this.Add(strconv.FormatUint(uint64(value), 10))
}

func (this *RPCParams) Build() string {
	length := len(this.Values)
	if length == 0 {
		return "[]"
	}

	result := "["
	for i := 0; i < length; i++ {
		result += this.Values[i]
		if i < (length - 1) {
			result += ", "
		}
	}
	result += "]"

	return result
}
