package examples

import "fmt"

func RunTransaction() {
	RunTransactionExecute()
	RunTransactionExecuteAndWatchFinalization()
	RunTransactionExecuteAndWatchInclusion()
	RunTransactionCustom()
	RunTransactionOptions()
	RunTransactionPayment()

	fmt.Println("RunTransaction finished correctly.")
}
