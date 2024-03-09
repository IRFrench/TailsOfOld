package search

import (
	"TailsOfOld/TailsOfOld/handlers/search"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/pocketbase"
)

func AddSearchRoutes(router *chi.Mux, database *pocketbase.PocketBase) {
	// Create index handler and route
	searchHandler := search.SearchHandler{Database: database}
	router.Handle("/search", searchHandler)
}
