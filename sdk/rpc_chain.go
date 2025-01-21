package sdk

import (
	prim "github.com/nmvalera/avail-go-sdk/primitives"
)

func (this *chainRPC) GetBlock(blockHash prim.Option[prim.H256]) (prim.Block, error) {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}
	value, err := this.client.Request("chain_getBlock", params.Build())
	if err != nil {
		return prim.Block{}, err
	}

	return prim.NewBlock(value)
}

func (this *chainRPC) GetBlockHash(blockNumber prim.Option[uint32]) (prim.H256, error) {
	params := RPCParams{}
	if blockNumber.IsSome() {
		params.AddUint32(blockNumber.Unwrap())
	}
	value, err := this.client.Request("chain_getBlockHash", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)
}

func (this *chainRPC) GetFinalizedHead() (prim.H256, error) {
	params := RPCParams{}
	value, err := this.client.Request("chain_getFinalizedHead", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)

}

func (this *chainRPC) GetHeader(blockHash prim.Option[prim.H256]) (prim.Header, error) {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}
	headerRaw, err := this.client.Request("chain_getHeader", params.Build())
	if err != nil {
		return prim.Header{}, err
	}

	return prim.NewHeaderFromJson(headerRaw)
}
