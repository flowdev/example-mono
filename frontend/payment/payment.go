package payment

import (
	"github.com/flowdev/example-mono/backend/catalogue"
	"github.com/flowdev/example-mono/email"
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/frontend/customer"
	"github.com/flowdev/example-mono/frontend/payment/banktransfer"
	"github.com/flowdev/example-mono/frontend/payment/creditcard"
	"github.com/flowdev/example-mono/frontend/payment/paypal"
	"github.com/flowdev/example-mono/frontend/payment/vouchers"
	"github.com/flowdev/example-mono/x/logging"
)

func PayFor(prodID, customerID, typ string) {
	var prod *entities.Product
	var cust *customer.Customer
	logging.Log("Pay for '" + prodID + "'")
	prod = catalogue.GetProduct(prodID)
	cust = customer.Login()

	switch typ {
	case "banktransfer":
		banktransfer.PayFor(prod)
	case "creditcard":
		creditcard.PayFor(prod)
	case "paypal":
		paypal.PayFor(prod)
	case "vouchers":
		vouchers.PayFor(prod)
	default:
		logging.Log("Unknown payment type '" + typ + "'")
	}
	email.Send(cust.Email, "Successful Purchase", "Congrats, ...")
}
