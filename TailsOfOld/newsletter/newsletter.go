package newsletter

import "net/smtp"

func SendNewsletter() error {
	auth := smtp.PlainAuth(
		"",
		"api",
		"fake_password",
		"live.smtp.mailtrap.io",
	)
	recipients := []string{
		"test@example.com",
	}
	message := []byte("To: recipient@example.net\r\n" +
		"From: mailtrap@demomailtrap.com\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	return smtp.SendMail("live.smtp.mailtrap.io:587", auth, "mailtrap@demomailtrap.com", recipients, message)
}
