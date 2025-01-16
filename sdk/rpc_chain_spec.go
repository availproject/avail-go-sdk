package sdk

import (
	prim "go-sdk/primitives"
)

func (this *chainSpecRPC) V1GenesisHash() (prim.H256, error) {
	params := RPCParams{}
	value, err := this.client.Request("chainSpec_v1_genesisHash", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)
}
