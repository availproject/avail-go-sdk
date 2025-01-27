package main

import (
	"github.com/availproject/avail-go-sdk/examples"
)

func main() {
	examples.RunBlockEvents()
}

func panic2(err error) {
	if err != nil {
		panic(err)
	}
}
