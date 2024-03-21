package articles

import (
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers/articles"
	"fmt"

	"github.com/go-chi/chi/v5"
)

func AddArticleRoutes(router *chi.Mux, database *db.DatabaseClient) {
	// Create articles overview handler and route

	gamesArticleHandler := articles.ArticleHandler{Database: database, Section: db.GAME_SECTION}
	router.Handle(fmt.Sprintf("/%v/{articleTitle}", db.GAME_SECTION), gamesArticleHandler)

	programmingArticleHandler := articles.ArticleHandler{Database: database, Section: db.PROGRAMMING_SECTION}
	router.Handle(fmt.Sprintf("/%v/{articleTitle}", db.PROGRAMMING_SECTION), programmingArticleHandler)

	tailsArticleHandler := articles.ArticleHandler{Database: database, Section: db.TALES_SECTION}
	router.Handle(fmt.Sprintf("/%v/{articleTitle}", db.TALES_SECTION), tailsArticleHandler)
}
