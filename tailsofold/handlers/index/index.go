package index

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers"
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"
	"html/template"
	"net/http"

	"github.com/rs/zerolog/log"
)

type IndexVariables struct {
	LatestGame        db.ArticleInfo
	LatestProgramming db.ArticleInfo
	LatestTale        db.ArticleInfo
}

type IndexHandler struct {
	Database *db.DatabaseClient
}

func (i IndexHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/index/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		weberrors.Borked(response, request)
		return
	}

	// Collect latest games article
	latestGamesArticle, err := i.Database.GetLatestSectionArticleInfo(db.GAME_SECTION)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest games article from database")
		weberrors.Borked(response, request)
		return
	}
	gamesAuthor, err := i.Database.GetAuthor(latestGamesArticle.Author)
	if err != nil {
		log.Error().Err(err).Str("game article title", latestGamesArticle.Title).Msg("failed to find game author")
		weberrors.Borked(response, request)
		return
	}
	latestGamesArticle.Author = gamesAuthor.Name

	// Collect latest programming article
	latestProgrammingArticle, err := i.Database.GetLatestSectionArticleInfo(db.PROGRAMMING_SECTION)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest programming article from database")
		weberrors.Borked(response, request)
		return
	}
	programmingAuthor, err := i.Database.GetAuthor(latestGamesArticle.Author)
	if err != nil {
		log.Error().Err(err).Str("programming article title", latestProgrammingArticle.Title).Msg("failed to find programming author")
		weberrors.Borked(response, request)
		return
	}
	latestProgrammingArticle.Author = programmingAuthor.Name

	// Collect latest tales article
	latestTalesArticle, err := i.Database.GetLatestSectionArticleInfo(db.TALES_SECTION)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest tales article from database")
		weberrors.Borked(response, request)
		return
	}
	talesAuthor, err := i.Database.GetAuthor(latestTalesArticle.Author)
	if err != nil {
		log.Error().Err(err).Str("tale article title", latestTalesArticle.Title).Msg("failed to find tale author")
		weberrors.Borked(response, request)
		return
	}
	latestTalesArticle.Author = talesAuthor.Name

	vars := IndexVariables{
		LatestProgramming: latestProgrammingArticle,
		LatestGame:        latestGamesArticle,
		LatestTale:        latestTalesArticle,
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
