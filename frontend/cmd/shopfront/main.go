package main

import (
	"github.com/flowdev/example-mono/frontend/customer"
	"github.com/flowdev/example-mono/frontend/payment"
	"github.com/flowdev/example-mono/frontend/shoppingcart"
	"github.com/flowdev/example-mono/x/logging"
)

func main() {
	logging.Log("Service started.")
	cust := customer.Login()
	logging.Log("Logged in " + cust.Email)
	payment.PayFor("Product123", "Customer123", "paypal")
	shoppingcart.AddProduct("Product123")
}
