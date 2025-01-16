package complex

import (
	prim "go-sdk/primitives"
)

func (this *chainSpecRPC) V1GenesisHash() prim.H256 {
	params := RPCParams{}
	value, err := this.client.Request("chainSpec_v1_genesisHash", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewH256FromHexString(value)
}
