package sdk

import (
	prim "github.com/availproject/avail-go-sdk/primitives"
)

type chainSpecRPC struct {
	client *Client
}

func (this *chainSpecRPC) V1GenesisHash() (prim.H256, error) {
	params := RPCParams{}
	value, err := this.client.RequestWithRetry("chainSpec_v1_genesisHash", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)
}
