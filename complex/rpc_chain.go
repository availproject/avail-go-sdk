package complex

import (
	prim "go-sdk/primitives"
)

func (this *chainRPC) GetBlock(blockHash prim.Option[prim.H256]) prim.Block {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}
	value, err := this.client.Request("chain_getBlock", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewBlock(value)
}

func (this *chainRPC) GetBlockHash(blockNumber prim.Option[uint32]) prim.H256 {
	params := RPCParams{}
	if blockNumber.IsSome() {
		params.AddUint32(blockNumber.Unwrap())
	}
	value, err := this.client.Request("chain_getBlockHash", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewH256FromHexString(value)
}

func (this *chainRPC) GetFinalizedHead() prim.H256 {
	params := RPCParams{}
	value, err := this.client.Request("chain_getFinalizedHead", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewH256FromHexString(value)

}

func (this *chainRPC) GetHeader(blockHash prim.Option[prim.H256]) prim.Header {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}
	headerRaw, err := this.client.Request("chain_getHeader", params.Build())
	if err != nil {
		panic(err)
	}

	return prim.NewHeaderFromJson(headerRaw)
}
