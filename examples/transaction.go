package examples

import "fmt"

func RunTransaction() {
	RunTransactionCustom()
	RunTransactionOptions()
	RunTransactionPayment()

	fmt.Println("RunTransaction finished correctly.")
}
