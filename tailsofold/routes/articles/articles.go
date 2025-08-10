package articles

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/articles"
	"fmt"

	"github.com/go-chi/chi/v5"
)

func AddArticleRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create articles overview handler and route

	programmingArticleHandler := articles.ArticleHandler{Database: database, Section: db.PROGRAMMING_SECTION}
	router.Handle(fmt.Sprintf("/%v/{articleTitle}", db.PROGRAMMING_SECTION), programmingArticleHandler)
}
