# Installation

You can find Go installation instructions [here](https://go.dev/doc/install). The required minimum version of Go is 1.23.0

If you already have installed Go, jump to -First Time Running- section.

## Installing GO in an Empty Ubuntu Container
Here are the instructions on how to install GO using the latest Ubuntu image.

```bash
podman run -it --rm --name ubuntu-container ubuntu:latest
apt-get update
apt-get upgrade
apt-get install nano wget
cd ./home/ubuntu/.
wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
tar -xf go1.23.5.linux-amd64.tar.gz
mv ./go /usr/local/go
export PATH=$PATH:/usr/local/go/bin
go version
# "go version go1.23.5 linux/amd64"
```


## First Time Running

1. Paste the following code to `main.go`:
```go
package main

import (
	"fmt"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk := SDK.NewSDK(SDK.TuringEndpoint)

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Alice
	acc, err := SDK.Account.Alice()
	if err != nil {
		panic(err)
	}

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}

	// Transaction Details
	println(fmt.Sprintf(`Block Hash: %v, Block Index: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))
}
```

2. Paste the following code to `go.mod`
```go
module mymodule

go 1.23.4

require github.com/availproject/avail-go-sdk v0.2.0-rc3
```

3. Fetch dependencies:
```bash
go mod tidy
```

4. Run Example:
```bash
go run .
```
