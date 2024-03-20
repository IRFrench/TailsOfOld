package newsletter

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"bytes"
	"fmt"
	"net/smtp"
	"net/url"
	"text/template"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/rs/zerolog/log"
)

type newsletterArgs struct {
	Articles  []db.ArticleInfo
	Recipient db.RecipientInfo
	Mailer    string
}

func (m *MailClient) SendNewsletter(database *pocketbase.PocketBase, since time.Time) error {
	// Collect newest articles
	newArticles, err := database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("created >= '%v' && %v = true", since.Format(time.DateTime), db.LIVE_COLUMN), "-created", 0, 0)
	if err != nil {
		log.Err(err).Msg("failed to find newest articles")
		return err
	}
	// Parse and populate article info
	allArticles := []db.ArticleInfo{}
	for _, article := range newArticles {
		author, err := database.Dao().FindRecordById(db.DB_USERS, article.GetString(db.AUTHOR_COLUMN))
		if err != nil {
			log.Err(err).Str("article title", article.GetString(db.TITLE_COLUMN)).Msg("failed to find author")
			return err
		}

		allArticles = append(allArticles, db.ArticleInfo{
			Title:       article.GetString(db.TITLE_COLUMN),
			Section:     article.GetString(db.SECTION_COLUMN),
			Description: article.GetString(db.DESCRIPTION_COLUMN),
			Author:      author.GetString(db.NAME_COLUMN),
			Created:     article.GetCreated().Time().Format(time.DateOnly),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", article.BaseFilesPath(), article.GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("www.tailsofold.com/%v/%v", article.GetString(db.SECTION_COLUMN), url.PathEscape(article.GetString(db.TITLE_COLUMN))),
		})
	}

	// If there are no new articles
	if len(allArticles) < 1 {
		log.Info().Msg("no new articles found this month")
		return nil
	}

	// Collect all recipients
	recipients, err := database.Dao().FindRecordsByExpr(db.DB_RECIPIENTS)
	if err != nil {
		log.Err(err).Msg("failed to find recipients")
		return err
	}
	// Parse recipients
	newsletterRecipients := []db.RecipientInfo{}
	for _, recipient := range recipients {
		newsletterRecipients = append(newsletterRecipients, db.RecipientInfo{
			FullName: recipient.GetString(db.FULL_NAME_COLUMN),
			Email:    recipient.Email(),
		})
	}

	// Send email
	for _, recipient := range newsletterRecipients {
		// Parse the template
		newsletter, err := template.New("newsletter").Parse(filesystem.Newsletter)
		if err != nil {
			return err
		}

		args := newsletterArgs{
			Articles:  allArticles,
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
