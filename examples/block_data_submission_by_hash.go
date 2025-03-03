package examples

import (
	"fmt"

	"github.com/availproject/avail-go-sdk/primitives"

	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func RunBlockDataSubmissionByHash() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	PanicOnError(err)

	blockHash, err := primitives.NewBlockHashFromHexString("0x94746ba186876d7407ee618d10cb6619befc59eeb173cacb00c14d1ff492fc58")
	PanicOnError(err)

	block, err := SDK.NewBlock(sdk.Client, blockHash)
	PanicOnError(err)

	// Block Blobs filtered by Transaction Hash
	txHash, err := primitives.NewH256FromHexString("0xe7efa71363d11bce370fe71a33e5ff296775f37507075c49316132131420f793")
	PanicOnError(err)

	blobs := block.DataSubmissions(SDK.Filter{}.WTxHash(txHash))
	AssertEq(len(blobs), 1, "")
	blob := &blobs[0]
	AssertEq(blob.TxHash.ToHuman(), txHash.ToHuman(), "Transaction Hash are not the same.")

	// Printout Block Blobs filtered by Transaction Hash
	accountId, err := primitives.NewAccountIdFromMultiAddress(blob.TxSigner)
	PanicOnError(err)
	fmt.Println(fmt.Sprintf(`Tx Hash: %v, Tx Index: %v, Data: %v, App Id: %v, Signer: %v,`, blob.TxHash, blob.TxIndex, string(blob.Data), blob.AppId, accountId.ToHuman()))

	fmt.Println("RunBlockDataSubmissionByHash finished correctly.")
}
