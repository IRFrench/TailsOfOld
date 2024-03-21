package search

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/search"

	"github.com/go-chi/chi/v5"
)

func AddSearchRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create index handler and route
	searchHandler := search.SearchHandler{Database: database}
	router.Handle("/search", searchHandler)
}
