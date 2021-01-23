package shoppingcart

import (
	"github.com/flowdev/example-mono/backend/catalogue"
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/x/logging"
)

func AddProduct(id string) {
	var prod *entities.Product
	logging.Log("AddProduct '" + id + "' to shopping cart")
	prod = catalogue.GetProduct(id)
	logging.Log("Added product " + prod.Name)
}
