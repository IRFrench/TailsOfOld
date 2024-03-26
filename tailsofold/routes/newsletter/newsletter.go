package newsletter

import (
	"TailsOfOld/internal/db"
	mailclient "TailsOfOld/internal/mail_client"
	"TailsOfOld/tailsofold/handlers/newsletter"

	"github.com/go-chi/chi/v5"
)

func AddNewsletterRoutes(router *chi.Mux, database *db.DatabaseClient, mail *mailclient.MailClient, newsletterFlag bool) {
	// Create newsletter subscribe handler and route
	subscribeHandler := newsletter.SubscribeHandler{Database: database, Mail: mail, Newsletter: newsletterFlag}
	router.Handle("/news/subscribe", subscribeHandler)

	if !newsletterFlag {
		return
	}
	// Create newsletter verify handler and route
	verifyHandler := newsletter.VerifyHandler{Database: database}
	router.Handle("/news/verify", verifyHandler)

	// Create newsletter unsubscribe handler and route
	unsubscribeHandler := newsletter.UnsubscribeHandler{Database: database}
	router.Handle("/news/unsubscribe", unsubscribeHandler)
}
