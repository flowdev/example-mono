package customer

import (
	"time"

	"github.com/flowdev/example-mono/email"
	"github.com/flowdev/example-mono/x/logging"
)

type Customer struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Name      string
}

func Login() *Customer {
	logging.Log("login user")
	return nil
}

func (c *Customer) ResetPassword() {
	email.Send(c.Email, "Reset your password", "bla...")
}
