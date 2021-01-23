package main

import (
	"github.com/flowdev/example-mono/backend/catalogue"
	"github.com/flowdev/example-mono/backend/recommendations/buyhist"
	"github.com/flowdev/example-mono/backend/recommendations/externalacts"
	"github.com/flowdev/example-mono/backend/recommendations/internalacts"
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/x/logging"
)

func main() {
	var prod *entities.Product
	prod = catalogue.GetProduct("Product123")
	logging.Log("Updating product " + prod.Name)

	buyhist.Recommend("Product123")
	internalacts.Recommend("Product123", "Customer123")
	externalacts.Recommend("Customer123")
}
