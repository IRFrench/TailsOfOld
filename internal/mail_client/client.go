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

func NewMailClient(config cfg.Configuration) *MailClient {
	auth := smtp.PlainAuth(
		"",
		config.Mail.Username,
		config.Mail.Password,
		config.Mail.Host,
	)
	return &MailClient{
		auth:   auth,
		host:   fmt.Sprintf("%v:587", config.Mail.Host),
		mailer: config.Mail.Mailer,
	}
}
