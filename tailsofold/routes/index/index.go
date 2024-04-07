package index

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/index"

	"github.com/go-chi/chi/v5"
)

func AddIndexRoutes(router *chi.Mux, database *db.DatabaseClient, newsletterSignup bool) {
	// Create index handler and route
	indexHandler := index.IndexHandler{Database: database, Newsletter: newsletterSignup}
	router.Handle("/", indexHandler)
}
