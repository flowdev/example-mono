package custstat

import (
	"github.com/flowdev/example-mono/frontend/customer"
	"github.com/flowdev/example-mono/x/logging"
)

func AddPurchase(custID string) {
	cust := customer.Login()
	logging.Log("Count purchase by '" + cust.Email + "'")
}
