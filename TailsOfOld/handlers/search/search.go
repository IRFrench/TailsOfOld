package search

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/handlers"
	weberrors "TailsOfOld/TailsOfOld/handlers/web_errors"
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
	"github.com/rs/zerolog/log"
)

type SearchHandler struct {
	Database *pocketbase.PocketBase
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

	articles, err := s.Database.Dao().FindRecordsByExpr(db.DB_ARTICLES, dbx.Like(db.TITLE_COLUMN, searchQuery), dbx.NewExp(fmt.Sprintf("%v = true", db.LIVE_COLUMN)))
	if err != nil {
		articles = []*models.Record{}
	}

	allArticles := []db.ArticleInfo{}

	for _, article := range articles {
		author, err := s.Database.Dao().FindRecordById(db.DB_USERS, article.GetString(db.AUTHOR_COLUMN))
		if err != nil {
			log.Error().Err(err).Str("article title", article.GetString(db.TITLE_COLUMN)).Msg("failed to find author")
			weberrors.Borked(response, request)
			return
		}

		allArticles = append(allArticles, db.ArticleInfo{
			Title:       article.GetString(db.TITLE_COLUMN),
			Section:     article.GetString(db.SECTION_COLUMN),
			Description: article.GetString(db.DESCRIPTION_COLUMN),
			Author:      author.GetString(db.NAME_COLUMN),
			Created:     article.GetCreated().Time().Format(time.DateOnly),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", article.BaseFilesPath(), article.GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", article.GetString(db.SECTION_COLUMN), url.PathEscape(article.GetString(db.TITLE_COLUMN))),
		})
	}

	vars := SearchVars{
		Articles: allArticles,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
