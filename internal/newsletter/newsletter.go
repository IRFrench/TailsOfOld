package newsletter

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
	"time"

	"github.com/rs/zerolog/log"
)

type newsletterArgs struct {
	Articles  []db.ArticleInfo
	Recipient db.RecipientInfo
	Mailer    string
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
		log.Info().Msg("no new articles found this month")
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
		// Parse the template
		newsletter, err := template.New("newsletter").Parse(filesystem.Newsletter)
		if err != nil {
			return err
		}

		args := newsletterArgs{
			Articles:  newArticles,
			Recipient: recipient,
			Mailer:    m.mailer,
		}

		var buffer bytes.Buffer
		if err := newsletter.Execute(&buffer, args); err != nil {
			return err
		}

		if err := smtp.SendMail(m.host, m.auth, m.mailer, []string{recipient.Email}, buffer.Bytes()); err != nil {
			return err
		}
	}
	return nil
}
