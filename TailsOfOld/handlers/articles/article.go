package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/handlers"
	weberrors "TailsOfOld/TailsOfOld/handlers/web_errors"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/rs/zerolog/log"
)

type ArticleHandler struct {
	Database *pocketbase.PocketBase
	Section  string
}

type ArticleVars struct {
	Article db.ArticleInfo
	Author  db.UserInfo
}

func (a ArticleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	articleTitle := chi.URLParam(request, "articleTitle")

	templatePath := "TailsOfOld/static/templates/articles/article.html"
	template := template.New("article.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	article, err := a.Database.Dao().FindFirstRecordByFilter(
		db.DB_ARTICLES,
		fmt.Sprintf("%v = {:title} && %v = '%v'", db.TITLE_COLUMN, db.SECTION_COLUMN, a.Section),
		dbx.Params{"title": articleTitle},
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to find article in database")
		weberrors.NotFoundHandler(response, request)
		return
	}

	author, err := a.Database.Dao().FindRecordById(db.DB_USERS, article.GetString(db.AUTHOR_COLUMN))
	if err != nil {
		log.Error().Err(err).Str("article title", article.GetString(db.TITLE_COLUMN)).Msg("failed to find article title")
		weberrors.Borked(response, request)
		return
	}

	vars := ArticleVars{
		Article: db.ArticleInfo{
			Title:     article.GetString(db.TITLE_COLUMN),
			Section:   article.GetString(db.SECTION_COLUMN),
			ImagePath: fmt.Sprintf("/pb_data/storage/%v/%v", article.BaseFilesPath(), article.GetString(db.IMAGEPATH_COLUMN)),
			Article:   article.GetString(db.ARTICLE_COLUMN),
			Created:   article.GetCreated().Time().Format(time.DateOnly),
			Updated:   article.GetUpdated().Time().Format(time.DateOnly),
		},
		Author: db.UserInfo{
			Name:       author.GetString(db.NAME_COLUMN),
			AvatarPath: fmt.Sprintf("/pb_data/storage/%v/%v", author.BaseFilesPath(), author.GetString(db.AVATAR_COLUMN)),
		},
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute template")
		weberrors.Borked(response, request)
		return
	}
}
