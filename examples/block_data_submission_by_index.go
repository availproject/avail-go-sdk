package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockDataSubmissionByIndex() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// Block Blobs filtered by Transaction Index
	blob := block.DataSubmissionByIndex(6).UnsafeUnwrap()
	AssertEq(blob.TxIndex, 6, "Transaction Indices are not the same.")

	// Printout Block Blobs filtered by Transaction Index
	accountId, err := metadata.NewAccountIdFromMultiAddress(blob.TxSigner)
	PanicOnError(err)
	fmt.Println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Data: %v, App Id: %v, Signer: %v,`, blob.TxHash, blob.TxIndex, string(blob.Data), blob.AppId, accountId.ToHuman()))

	fmt.Println("RunBlockDataSubmissionByIndex finished correctly.")
}
