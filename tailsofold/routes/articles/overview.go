package articles

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/articles"
	"fmt"

	"github.com/go-chi/chi/v5"
)

func AddArticleOverviewRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create programming overview handler and route
	programmingOverviewHandler := articles.OverviewHandler{Database: database, Section: db.PROGRAMMING_SECTION}
	router.Handle(fmt.Sprintf("/%v", db.PROGRAMMING_SECTION), programmingOverviewHandler)
}
