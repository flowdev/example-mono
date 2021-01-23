package internalacts

import (
	"github.com/flowdev/example-mono/backend/catalogue"
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/frontend/customer"
	"github.com/flowdev/example-mono/x/logging"
)

func Recommend(prodID, customerID string) {
	var prod *entities.Product
	var cust *customer.Customer
	prod = catalogue.GetProduct(prodID)
	logging.Log("Recommend " + prod.Name)
	cust = customer.Login()
	logging.Log("Recommend for " + cust.Email)
}
