package examples

import (
	"fmt"

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
		fmt.Println("Block Number:", value.Header.Number)
	}

	{
		// chain_GetBlockHash
		value, err := sdk.Client.Rpc.Chain.GetBlockHash(prim.NewNone[uint32]())
		PanicOnError(err)
		fmt.Println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetFinalizedHead
		value, err := sdk.Client.Rpc.Chain.GetFinalizedHead()
		PanicOnError(err)
		fmt.Println("Block Hash:", value.ToHuman())
	}

	{
		// chain_GetHeader
		value, err := sdk.Client.Rpc.Chain.GetHeader(prim.NewNone[prim.H256]())
		PanicOnError(err)
		fmt.Println("Block Number:", value.Number)
	}

	{
		// chainspec_V1GenesisHash
		value, err := sdk.Client.Rpc.ChainSpec.V1GenesisHash()
		PanicOnError(err)
		fmt.Println("Genesis Hash:", value.ToHuman())
	}

	{
		// system_AccountNextIndex
		value, err := sdk.Client.Rpc.System.AccountNextIndex("5GEQ6S3vpSFjYCqsrndQhcPL3sh8uAYbpeCiZFhF4u9EjK6F")
		PanicOnError(err)
		fmt.Println("Nonce:", value)
	}

	{
		// system_Chain
		value, err := sdk.Client.Rpc.System.Chain()
		PanicOnError(err)
		fmt.Println("Chain:", value)
	}

	{
		// system_ChainType
		value, err := sdk.Client.Rpc.System.ChainType()
		PanicOnError(err)
		fmt.Println("ChainType:", value)
	}

	{
		// system_Health
		value, err := sdk.Client.Rpc.System.Health()
		PanicOnError(err)
		fmt.Println("Health: IsSyncing:", value.IsSyncing)
	}

	{
		// system_LocalPeerId
		value, err := sdk.Client.Rpc.System.LocalPeerId()
		PanicOnError(err)
		fmt.Println("Local Peer Id:", value)
	}

	{
		// system_Name
		value, err := sdk.Client.Rpc.System.Name()
		PanicOnError(err)
		fmt.Println("Name:", value)
	}

	{
		// system_NodeRoles
		value, err := sdk.Client.Rpc.System.NodeRoles()
		PanicOnError(err)
		for _, elem := range value {
			fmt.Println("Role:", elem)
		}
	}

	{
		// system_Properties
		value, err := sdk.Client.Rpc.System.Properties()
		PanicOnError(err)
		fmt.Println("Ss58format:", value.Ss58Format)
		fmt.Println("Token Symbol:", value.TokenSymbol)
	}

	{
		// system_SyncState
		value, err := sdk.Client.Rpc.System.SyncState()
		PanicOnError(err)
		fmt.Println("Starting Block:", value.StartingBlock)
		fmt.Println("Current Block:", value.CurrentBlock)
		fmt.Println("Highest Block:", value.HighestBlock)
	}

	{
		// system_Version
		value, err := sdk.Client.Rpc.System.Version()
		PanicOnError(err)
		fmt.Println("Version:", value)
	}

	{
		// author_RotateKeys
		_, _ = sdk.Client.Rpc.Author.RotateKeys()
	}

	fmt.Println("RunRpc finished correctly.")
}
