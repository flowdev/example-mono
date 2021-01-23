package externalacts

import (
	"github.com/flowdev/example-mono/frontend/customer"
	"github.com/flowdev/example-mono/x/logging"
)

func Recommend(custID string) {
	cust := customer.Login()
	logging.Log("Recommend for '" + cust.Email + "'")
}
