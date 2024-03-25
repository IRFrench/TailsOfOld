package mailclient

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"bytes"
	"fmt"
	"net/smtp"
	"net/url"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"
)

type newsletterArgs struct {
	Articles        []db.ArticleInfo
	Recipient       db.RecipientInfo
	Mailer          string
	UnsubscribePath string
}

func (m *MailClient) SendNewsletter(database *db.DatabaseClient, since time.Time) error {
	// Collect newest articles
	newArticles, err := database.GetArticlesCreatedSinceTime(since)
	if err != nil {
		log.Err(err).Msg("failed to find newest articles")
		return err
	}
	// Parse and populate article info
	for index := range newArticles {
		author, err := database.GetAuthor(newArticles[index].Author)
		if err != nil {
			log.Err(err).Str("article title", newArticles[index].Title).Msg("failed to find author")
			return err
		}
		newArticles[index].Author = author.Name
		newArticles[index].ArticlePath = fmt.Sprintf("%v%v", "www.tailsofold.com", newArticles[index].ArticlePath)
		newArticles[index].ImagePath = fmt.Sprintf("%v%v", "www.tailsofold.com", newArticles[index].ImagePath)
	}

	// If there are no new articles
	if len(newArticles) < 1 {
		log.Debug().Msg("no new articles found this month")
		return nil
	}

	// Collect all recipients
	recipients, err := database.GetVerifiedRecipients()
	if err != nil {
		log.Err(err).Msg("failed to find recipients")
		return err
	}

	// Send email
	for _, recipient := range recipients {
		log.Debug().Str("email", recipient.Email).Msg("building newsletter")

		// Parse the template
		newsletter, err := template.New("newsletter").Parse(filesystem.Newsletter)
		if err != nil {
			log.Err(err).Msg("failed to parse template")
			return err
		}

		args := newsletterArgs{
			Articles:        newArticles,
			Recipient:       recipient,
			Mailer:          m.mailer,
			UnsubscribePath: fmt.Sprintf("www.tailsofold.com/news/unsubscribe?id=%v&email=%v", url.QueryEscape(recipient.Id), url.QueryEscape(recipient.Email)),
		}

		var buffer bytes.Buffer
		if err := newsletter.Execute(&buffer, args); err != nil {
			log.Err(err).Msg("failed to execute template")
			return err
		}

		if err := smtp.SendMail(m.host, m.auth, m.mailer, []string{recipient.Email}, buffer.Bytes()); err != nil {
			log.Err(err).Msg("failed to send newsletter")
			return err
		}
		log.Debug().Str("email", recipient.Email).Msg("newsletter sent")
	}
	return nil
}

type VerifyArgs struct {
	Recipient       db.RecipientInfo
	Mailer          string
	VerifyPath      string
	UnsubscribePath string
}

func (m *MailClient) SendVerification(database *db.DatabaseClient, recipient db.RecipientInfo) error {
	// Parse the template
	newsletter, err := template.New("verify").Parse(filesystem.Verify)
	if err != nil {
		return err
	}

	args := VerifyArgs{
		Recipient:       recipient,
		Mailer:          m.mailer,
		VerifyPath:      fmt.Sprintf("www.tailsofold.com/news/verify?id=%v&email=%v", url.QueryEscape(recipient.Id), url.QueryEscape(recipient.Email)),
		UnsubscribePath: fmt.Sprintf("www.tailsofold.com/news/unsubscribe?id=%v&email=%v", url.QueryEscape(recipient.Id), url.QueryEscape(recipient.Email)),
	}

	var buffer bytes.Buffer
	if err := newsletter.Execute(&buffer, args); err != nil {
		return err
	}

	if err := smtp.SendMail(m.host, m.auth, m.mailer, []string{recipient.Email}, buffer.Bytes()); err != nil {
		return err
	}
	return nil
}
