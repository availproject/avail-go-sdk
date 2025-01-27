package main

import (
	"github.com/availproject/avail-go-sdk/examples"
)

func main() {
	examples.RunBlockTransactionAll()
	examples.RunBlockTransactionBySigner()
	examples.RunBlockTransactionByIndex()
	examples.RunBlockTransactionByHash()
	examples.RunBlockTransactionByAppId()

}

func panic2(err error) {
	if err != nil {
		panic(err)
	}
}
