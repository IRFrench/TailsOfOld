package articles

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/articles"
	"fmt"

	"github.com/go-chi/chi/v5"
)

func AddArticleOverviewRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create games overview handler and route
	gamesOverviewHandler := articles.OverviewHandler{Database: database, Section: db.GAME_SECTION}
	router.Handle(fmt.Sprintf("/%v", db.GAME_SECTION), gamesOverviewHandler)

	// Create programming overview handler and route
	programmingOverviewHandler := articles.OverviewHandler{Database: database, Section: db.PROGRAMMING_SECTION}
	router.Handle(fmt.Sprintf("/%v", db.PROGRAMMING_SECTION), programmingOverviewHandler)

	// Create tales overview handler and route
	talesOverviewHandler := articles.OverviewHandler{Database: database, Section: db.TALES_SECTION}
	router.Handle(fmt.Sprintf("/%v", db.TALES_SECTION), talesOverviewHandler)
}
