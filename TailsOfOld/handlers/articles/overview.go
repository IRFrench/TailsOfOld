package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/handlers"
	weberrors "TailsOfOld/TailsOfOld/handlers/web_errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/rs/zerolog/log"
)

type OverviewHandler struct {
	Database *pocketbase.PocketBase
	Section  string
}

type OverviewVariables struct {
	Section  string
	Articles []db.ArticleInfo
}

func (o OverviewHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/articles/overview.html"
	template := template.New("overview.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into template")
		weberrors.Borked(response, request)
		return
	}

	// Gather articles
	databaseArticles, err := o.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, o.Section, db.LIVE_COLUMN), "-created", 0, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to find section articles in the database")
		weberrors.Borked(response, request)
		return
	}

	allArticles := []db.ArticleInfo{}

	for _, article := range databaseArticles {
		articleAuthor, err := o.Database.Dao().FindRecordById(db.DB_USERS, article.GetString(db.AUTHOR_COLUMN))
		if err != nil {
			log.Error().Err(err).Str("article title", article.GetString(db.TITLE_COLUMN)).Msg("failed to find article author")
			weberrors.Borked(response, request)
			return
		}

		newFlag := false
		articleCreationDate := article.GetCreated().Time()
		if time.Now().Before(articleCreationDate.Add(24 * time.Hour)) {
			newFlag = true
		}

		allArticles = append(allArticles, db.ArticleInfo{
			Title:       article.GetString(db.TITLE_COLUMN),
			Description: article.GetString(db.DESCRIPTION_COLUMN),
			Created:     article.GetCreated().Time().Format(time.DateOnly),
			Author:      articleAuthor.GetString(db.NAME_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", article.BaseFilesPath(), article.GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", o.Section, url.PathEscape(article.GetString(db.TITLE_COLUMN))),
			New:         newFlag,
		})
	}

	vars := OverviewVariables{
		Section:  o.Section,
		Articles: allArticles,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
