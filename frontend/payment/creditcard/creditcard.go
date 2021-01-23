package creditcard

import (
	"github.com/flowdev/example-mono/entities"
	"github.com/flowdev/example-mono/x/logging"
)

func PayFor(prod *entities.Product) {
	logging.Log("Added product " + prod.Name)
}
