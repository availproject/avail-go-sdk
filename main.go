package main

import "github.com/availproject/avail-go-sdk/examples"

func main() {
	examples.RunBlockTransactions()

}

func panic2(err error) {
	if err != nil {
		panic(err)
	}
}
