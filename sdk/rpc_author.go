package sdk

import (
	"strings"

	prim "github.com/availproject/avail-go-sdk/primitives"
)

type authorRPC struct {
	client *Client
}

func (this *authorRPC) RotateKeys() (string, error) {
	params := RPCParams{}
	return this.client.RequestWithRetry("author_rotateKeys", params.Build())
}

// Transaction needs to be hex and scale encoded
func (this *authorRPC) SubmitExtrinsic(tx string) (prim.H256, error) {
	if !strings.HasPrefix(tx, "0x") {
		tx = "0x" + tx
	}
	params := RPCParams{}
	params.Add("\"" + tx + "\"")

	txHash, err := this.client.RequestWithRetry("author_submitExtrinsic", params.Build())
	if err != nil {
		return prim.H256{}, err
	}

	return prim.NewH256FromHexString(txHash)
}
