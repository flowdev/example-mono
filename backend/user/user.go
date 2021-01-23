package user

import (
	"time"

	"github.com/flowdev/example-mono/email"
	"github.com/flowdev/example-mono/x/logging"
)

type User struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Name      string
}

func Login() *User {
	logging.Log("login user")
	return nil
}

func (c *User) ResetPassword() {
	email.Send(c.Email, "Reset your password", "bla...")
}
