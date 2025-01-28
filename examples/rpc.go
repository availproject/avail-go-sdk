package examples

import (
	prim "github.com/availproject/avail-go-sdk/primitives"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunRpc() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	{
		// chain_GetBlock
		value, err := sdk.Client.Rpc.Chain.GetBlock(prim.NewNone[prim.H256]())
		PanicOnError(err)
		println("Block Number:", value.Header.Number)
	}

	{
		// chain_GetBlockHash
		value, err := sdk.Client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
		PanicOnError(err)
		println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetFinalizedHead
		value, err := sdk.Client.Rpc.Chain.GetFinalizedHead()
		PanicOnError(err)
		println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetHeader
		value, err := sdk.Client.Rpc.Chain.GetHeader(prim.NewNone[prim.H256]())
		PanicOnError(err)
		println("Block Number:", value.Number)
	}

	{
		// chainspec_V1GenesisHash
		value, err := sdk.Client.Rpc.ChainSpec.V1GenesisHash()
		PanicOnError(err)
		println("Genesis Hash:", value.ToHuman())
	}

	{
		// system_AccountNextIndex
		value, err := sdk.Client.Rpc.System.AccountNextIndex("5GEQ6S3vpSFjYCqsrndQhcPL3sh8uAYbpeCiZFhF4u9EjK6F")
		PanicOnError(err)
		println("Nonce:", value)
	}

	{
		// system_Chain
		value, err := sdk.Client.Rpc.System.Chain()
		PanicOnError(err)
		println("Chain:", value)
	}

	{
		// system_ChainType
		value, err := sdk.Client.Rpc.System.ChainType()
		PanicOnError(err)
		println("ChainType:", value)
	}

	{
		// system_Health
		value, err := sdk.Client.Rpc.System.Health()
		PanicOnError(err)
		println("Health: IsSyncing:", value.IsSyncing)
	}

	{
		// system_LocalPeerId
		value, err := sdk.Client.Rpc.System.LocalPeerId()
		PanicOnError(err)
		println("Local Peer Id:", value)
	}

	{
		// system_Name
		value, err := sdk.Client.Rpc.System.Name()
		PanicOnError(err)
		println("Name:", value)
	}

	{
		// system_NodeRoles
		value, err := sdk.Client.Rpc.System.NodeRoles()
		PanicOnError(err)
		for _, elem := range value {
			println("Role:", elem)
		}
	}

	{
		// system_Properties
		value, err := sdk.Client.Rpc.System.Properties()
		PanicOnError(err)
		println("Ss58format:", value.Ss58Format)
		println("Token Symbol:", value.TokenSymbol)
	}

	{
		// system_SyncState
		value, err := sdk.Client.Rpc.System.SyncState()
		PanicOnError(err)
		println("Current Block:", value.CurrentBlock)
		println("Highest Block:", value.HighestBlock)
	}

	{
		// system_Version
		value, err := sdk.Client.Rpc.System.Version()
		PanicOnError(err)
		println("Version:", value)
	}

	{
		// author_RotateKeys
		_, _ = sdk.Client.Rpc.Author.RotateKeys()
	}

	println("RunRpc finished correctly.")
}
