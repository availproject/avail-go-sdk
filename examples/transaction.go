package examples

import "fmt"

func RunTransaction() {
	RunTransactionExecute()
	RunTransactionExecuteAndWatch()
	RunTransactionExecuteAndWatchFinalization()
	RunTransactionExecuteAndWatchInclusion()
	RunTransactionCustom()
	RunTransactionOptions()
	RunTransactionPayment()

	fmt.Println("RunTransaction finished correctly.")
}
