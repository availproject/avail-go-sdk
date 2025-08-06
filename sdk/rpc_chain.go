package sdk

import (
	"fmt"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type chainRPC struct {
	client *Client
}

func (c *chainRPC) GetBlock(blockHash prim.Option[prim.H256]) (prim.Block, error) {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	value, err := c.client.RequestWithRetry("chain_getBlock", params.Build())

	if err != nil {
		fmt.Println(fmt.Sprintf("Value: %v, Error: %v", value, err))
		return prim.Block{}, err
	}

	return prim.NewBlock(value)
}

func (c *chainRPC) GetBlockHash(blockNumber prim.Option[uint32]) (prim.H256, error) {
	params := RPCParams{}
	if blockNumber.IsSome() {
		params.AddUint32(blockNumber.Unwrap())
	}

	value, err := c.client.RequestWithRetry("chain_getBlockHash", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)
}

func (c *chainRPC) GetFinalizedHead() (prim.H256, error) {
	params := RPCParams{}

	value, err := c.client.RequestWithRetry("chain_getFinalizedHead", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(value)

}

func (c *chainRPC) GetHeader(blockHash prim.Option[prim.H256]) (prim.Header, error) {
	params := RPCParams{}
	if blockHash.IsSome() {
		params.AddH256(blockHash.Unwrap())
	}

	value, err := c.client.RequestWithRetry("chain_getHeader", params.Build())
	if err != nil {
		return prim.Header{}, err
	}

	return prim.NewHeaderFromJson(value)
}
