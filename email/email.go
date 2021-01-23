package email

import "github.com/flowdev/example-mono/x/logging"

func Send(address, title, text string) {
	logging.Log("Sending email to: " + address)
}
