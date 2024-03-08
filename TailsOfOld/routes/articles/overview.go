package articles

import (
	"TailsOfOld/TailsOfOld/handlers/articles"
	"TailsOfOld/TailsOfOld/section"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/pocketbase"
)

func AddArticleRoutes(router *chi.Mux, database *pocketbase.PocketBase) {
	// Create games overview handler and route
	gamesOverviewHandler := articles.OverviewHandler{Database: database, Section: section.GAME_SECTION}
	router.Handle(fmt.Sprintf("/%v", section.GAME_SECTION), gamesOverviewHandler)

	// Create programming overview handler and route
	programmingOverviewHandler := articles.OverviewHandler{Database: database, Section: section.PROGRAMMING_SECTION}
	router.Handle(fmt.Sprintf("/%v", section.PROGRAMMING_SECTION), programmingOverviewHandler)

	// Create tales overview handler and route
	talesOverviewHandler := articles.OverviewHandler{Database: database, Section: section.TALES_SECTION}
	router.Handle(fmt.Sprintf("/%v", section.TALES_SECTION), talesOverviewHandler)
}
