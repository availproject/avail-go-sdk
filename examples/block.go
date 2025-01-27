package examples

import ()

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

	println("RunBlock finished correctly.")
}
