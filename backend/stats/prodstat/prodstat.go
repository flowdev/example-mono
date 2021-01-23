package prodstat

import (
	"github.com/flowdev/example-mono/backend/catalogue"
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/x/logging"
)

func AddPurchase(prodID string) {
	var prod *entities.Product
	prod = catalogue.GetProduct(prodID)
	logging.Log("Count purchase of " + prod.Name)
}
