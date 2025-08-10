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
	LatestArticles []db.ArticleInfo
}

type IndexHandler struct {
	Database *db.DatabaseClient
}

func (i IndexHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "tailsofold/static/templates/index/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		weberrors.Borked(response, request)
		return
	}

	// Collect latest programming article
	latestArticles, err := i.Database.GetLatestArticlesInfo()
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest programming article from database")
		weberrors.Borked(response, request)
		return
	}
	for index := range latestArticles {
		articleAuthor, err := i.Database.GetAuthor(latestArticles[index].Author)
		if err != nil {
			log.Error().Err(err).Str("programming article title", latestArticles[index].Title).Msg("failed to find programming author")
			weberrors.Borked(response, request)
			return
		}
		latestArticles[index].Author = articleAuthor.Name
	}

	vars := IndexVariables{
		LatestArticles: latestArticles,
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
