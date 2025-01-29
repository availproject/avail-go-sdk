package examples

import "fmt"

func RunBlock() {
	RunBlockTransactionAll()
	RunBlockTransactionByAppId()
	RunBlockTransactionByHash()
	RunBlockTransactionByIndex()
	RunBlockTransactionBySigner()
	RunBlockDataSubmissionAll()
	RunBlockDataSubmissionByAppId()
	RunBlockDataSubmissionByHash()
	RunBlockDataSubmissionByIndex()
	RunBlockDataSubmissionBySigner()
	RunBlockEvents()

	fmt.Println("RunBlock finished correctly.")
}
