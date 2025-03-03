package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockDataSubmissionByAppId() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// Block Blobs filtered by App Id
	appId := uint32(2)
	blobs := block.DataSubmissions(SDK.Filter{}.WAppId(appId))
	AssertEq(len(blobs), 2, "Data Submission count is not 2")

	// Printout Block Blobs filtered by App Id
	for _, blob := range blobs {
		AssertEq(blob.AppId, appId, "Transaction App Ids are not the same.")

		accountId, err := primitives.NewAccountIdFromMultiAddress(blob.TxSigner)
		PanicOnError(err)

		fmt.Println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Data: %v, App Id: %v, Signer: %v,`, blob.TxHash, blob.TxIndex, string(blob.Data), blob.AppId, accountId.ToHuman()))
	}

	fmt.Println("RunBlockDataSubmissionByAppId finished correctly.")
}
