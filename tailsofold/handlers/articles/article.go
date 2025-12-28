package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/handlers"
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"
	"net/http"
	"path"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type ArticleHandler struct {
	Database *db.DatabaseClient
	Section  string
}

type ArticleVars struct {
	Article db.ArticleInfo
	Author  db.UserInfo
}

func (a ArticleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	articleTitle := chi.URLParam(request, "articleTitle")

	templatePath := "tailsofold/static/templates/articles/article.html"
	template := template.New("article.html")

	article, err := a.Database.GetFullArticle(articleTitle, a.Section)
	if err != nil {
		log.Error().Err(err).Msg("failed to find article in database")
		weberrors.NotFoundHandler(response, request)
		return
	}

	insertPath := getInsertPath(article.Insert)

	template, err = template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath, insertPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	author, err := a.Database.GetAuthor(article.Author)
	if err != nil {
		log.Error().Err(err).Str("article title", article.Title).Msg("failed to find article title")
		weberrors.Borked(response, request)
		return
	}

	vars := ArticleVars{
		Article: article,
		Author:  author,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute template")
		weberrors.Borked(response, request)
		return
	}
}

func getInsertPath(insertName string) string {
	entires, _ := filesystem.FileSystem.ReadDir(handlers.INSERT_TEMPLATES)
	for index := range entires {
		if entires[index].Name() == insertName {
			return path.Join(handlers.INSERT_TEMPLATES, insertName)
		}
	}
	return path.Join(handlers.INSERT_TEMPLATES, "empty.html")
}
