package mailclient

import (
	"TailsOfOld/cfg"
	"fmt"
	"net/smtp"
)

type MailClient struct {
	auth     smtp.Auth
	host     string
	mailer   string
	template string
}

func NewMailClient(config cfg.Mail) *MailClient {
	auth := smtp.PlainAuth(
		"",
		config.Username,
		config.Password,
		config.Host,
	)
	return &MailClient{
		auth:   auth,
		host:   fmt.Sprintf("%v:587", config.Host),
		mailer: config.Mailer,
	}
}
