package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/metadata"
	"github.com/availproject/avail-go-sdk/primitives"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockDataSubmissionBySigner() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	accountId, err := metadata.NewAccountIdFromAddress("5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty")
	PanicOnError(err)

	// Block Blobs filtered by Signer
	blobs := block.DataSubmissionBySigner(accountId)
	println("Blob Count: ", len(blobs))
	AssertEq(len(blobs), 1, "Data Submission count is not 1")

	// Printout Block Blobs filtered by Signer
	for _, blob := range blobs {
		blobAccountId, err := metadata.NewAccountIdFromMultiAddress(blob.TxSigner)
		PanicOnError(err)
		AssertEq(blobAccountId.ToHuman(), accountId.ToHuman(), "Transaction Signers are not the same.")

		println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Data: %v, App Id: %v, Signer: %v,`, blob.TxHash, blob.TxIndex, string(blob.Data), blob.AppId, blobAccountId.ToHuman()))
	}

	println("RunBlockDataSubmissionBySigner finished correctly.")
}
