package index

import (
	"TailsOfOld/TailsOfOld/handlers/index"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/pocketbase"
)

func AddIndexRoutes(router *chi.Mux, database *pocketbase.PocketBase) {
	// Create index handler and route
	indexHandler := index.IndexHandler{Database: database}
	router.Handle("/", indexHandler)
}
