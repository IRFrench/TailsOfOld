package index

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/index"

	"github.com/go-chi/chi/v5"
)

func AddIndexRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create index handler and route
	indexHandler := index.IndexHandler{Database: database}
	router.Handle("/", indexHandler)
}
