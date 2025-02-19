# Installation

You can find Go installation instructions [here](https://go.dev/doc/install). The required minimum version of Go is 1.23.0

If you already have installed Go (version no less than 1.23.0), jump to `Add Avail-GO SDK as dependency`  section.

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
## Add Avail-GO SDK as dependency 

#### To Existing Project

```bash
# Fetches Avail-GO SDK v0.2.2. This might not be the newest version so make sure to check out the latest github avail-go-sdk release.
# Link to Github: https://github.com/availproject/avail-go-sdk/releases
go get github.com/availproject/avail-go-sdk@v0.2.2
```

#### To A New Project

```bash
# Creates a new project with name myproject
go mod init myproject
# Fetches Avail-GO SDK v0.2.2. This might not be the newest version so make sure to check out the latest github avail-go-sdk release.
# Link to Github: https://github.com/availproject/avail-go-sdk/releases
go get github.com/availproject/avail-go-sdk@v0.2.2
```

#### First Time Running

1. Paste the following code to `main.go`:
```go
package main

import (
	"fmt"
	SDK "github.com/availproject/avail-go-sdk/sdk"
)

func main() {
	sdk, err := SDK.NewSDK(SDK.TuringEndpoint)
	if err != nil {
		panic(err)
	}

	// Use SDK.Account.NewKeyPair("Your key") to use a different account than Ferdie
	acc := SDK.Account.Ferdie()

	tx := sdk.Tx.DataAvailability.SubmitData([]byte("MyData"))
	fmt.Println("Submitting new Transaction... Can take up to 20 seconds")
	res, err := tx.ExecuteAndWatchInclusion(acc, SDK.NewTransactionOptions().WithAppId(1))
	if err != nil {
		panic(err)
	}

	// Transaction Details
	fmt.Println(fmt.Sprintf(`Block Hash: %v, Block Number: %v, Tx Hash: %v, Tx Index: %v`, res.BlockHash.ToHexWith0x(), res.BlockNumber, res.TxHash.ToHexWith0x(), res.TxIndex))
}
```

2. Fetch dependencies:
```bash
go mod tidy
```

3. Run Example:
```bash
go run .
```
