package examples

import (
	prim "github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunRpc() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	{
		// chain_GetBlock
		value, err := sdk.Client.Rpc.Chain.GetBlock(prim.NewNone[prim.H256]())
		if err != nil {
			panic(err)
		}
		println("Block Number:", value.Header.Number)
	}

	{
		// chain_GetBlockHash
		value, err := sdk.Client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
		if err != nil {
			panic(err)
		}
		println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetFinalizedHead
		value, err := sdk.Client.Rpc.Chain.GetFinalizedHead()
		if err != nil {
			panic(err)
		}
		println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetHeader
		value, err := sdk.Client.Rpc.Chain.GetHeader(prim.NewNone[prim.H256]())
		if err != nil {
			panic(err)
		}
		println("Block Number:", value.Number)
	}

	{
		// chainspec_V1GenesisHash
		value, err := sdk.Client.Rpc.ChainSpec.V1GenesisHash()
		if err != nil {
			panic(err)
		}
		println("Genesis Hash:", value.ToHuman())
	}

	{
		// system_AccountNextIndex
		value, err := sdk.Client.Rpc.System.AccountNextIndex("5GEQ6S3vpSFjYCqsrndQhcPL3sh8uAYbpeCiZFhF4u9EjK6F")
		if err != nil {
			panic(err)
		}
		println("Nonce:", value)
	}

	{
		// system_Chain
		value, err := sdk.Client.Rpc.System.Chain()
		if err != nil {
			panic(err)
		}
		println("Chain:", value)
	}

	{
		// system_ChainType
		value, err := sdk.Client.Rpc.System.ChainType()
		if err != nil {
			panic(err)
		}
		println("ChainType:", value)
	}

	{
		// system_Health
		value, err := sdk.Client.Rpc.System.Health()
		if err != nil {
			panic(err)
		}
		println("Health: IsSyncing:", value.IsSyncing)
	}

	{
		// system_LocalPeerId
		value, err := sdk.Client.Rpc.System.LocalPeerId()
		if err != nil {
			panic(err)
		}
		println("Local Peer Id:", value)
	}

	{
		// system_Name
		value, err := sdk.Client.Rpc.System.Name()
		if err != nil {
			panic(err)
		}
		println("Name:", value)
	}

	{
		// system_NodeRoles
		value, err := sdk.Client.Rpc.System.NodeRoles()
		if err != nil {
			panic(err)
		}
		for _, elem := range value {
			println("Role:", elem)
		}
	}

	{
		// system_Properties
		value, err := sdk.Client.Rpc.System.Properties()
		if err != nil {
			panic(err)
		}
		println("Ss58format:", value.Ss58Format)
		println("Token Symbol:", value.TokenSymbol)
	}

	{
		// system_SyncState
		value, err := sdk.Client.Rpc.System.SyncState()
		if err != nil {
			panic(err)
		}
		println("Current Block:", value.CurrentBlock)
		println("Highest Block:", value.HighestBlock)
	}

	{
		// system_Version
		value, err := sdk.Client.Rpc.System.Version()
		if err != nil {
			panic(err)
		}
		println("Version:", value)
	}

	{
		// author_RotateKeys
		_, _ = sdk.Client.Rpc.Author.RotateKeys()
	}

	println("RunRpc finished correctly.")
}
