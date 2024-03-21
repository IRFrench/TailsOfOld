package search

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers"
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

type SearchHandler struct {
	Database *db.DatabaseClient
}

type SearchVars struct {
	Articles []db.ArticleInfo
}

func (s SearchHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/search/search.html"
	template := template.New("search.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		weberrors.Borked(response, request)
		return
	}

	searchQuery := request.URL.Query().Get("q")

	articles, err := s.Database.SearchArticlesByTitle(searchQuery)
	if err != nil {
		log.Error().Err(err).Msg("failed to search for articles with that title")
		weberrors.Borked(response, request)
		return
	}

	for _, article := range articles {
		author, err := s.Database.GetAuthor(article.Author)
		if err != nil {
			log.Error().Err(err).Str("article title", article.Title).Msg("failed to find author")
			weberrors.Borked(response, request)
			return
		}
		article.Author = author.Name
	}

	vars := SearchVars{
		Articles: articles,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
