package catalogue

import (
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/x/logging"
)

func GetProduct(id string) *entities.Product {
	logging.Log("GetProduct: " + id)
	return nil
}
