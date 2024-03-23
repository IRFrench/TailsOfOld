package newsletter

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/newsletter"

	"github.com/go-chi/chi/v5"
)

func AddNewsletterRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create newsletter subscribe handler and route
	subscribeHandler := newsletter.SubscribeHandler{Database: database}
	router.Handle("/news/subscribe", subscribeHandler)

	// Create newsletter verify handler and route
	verifyHandler := newsletter.VerifyHandler{Database: database}
	router.Handle("/news/verify", verifyHandler)

	// Create newsletter unsubscribe handler and route
	unsubscribeHandler := newsletter.UnsubscribeHandler{Database: database}
	router.Handle("/news/unsubscribe", unsubscribeHandler)
}
