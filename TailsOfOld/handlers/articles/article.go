package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/handlers"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/pocketbase"
)

type ArticleHandler struct {
	Database *pocketbase.PocketBase
	Section  string
}

type ArticleVariables struct {
	Title     string
	Author    string
	Section   string
	ImagePath string
	Article   string
	Created   string
	Updated   string
}

func (a ArticleHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	articleTitle := chi.URLParam(request, "articleTitle")

	templatePath := "TailsOfOld/static/templates/articles/article.html"
	template := template.New("article.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		slog.Error("Failed to parse file system into template", err)
		panic(err)
	}

	article, err := a.Database.Dao().FindFirstRecordByFilter(
		db.DB_ARTICLES,
		fmt.Sprintf("%v = '%v' && %v = '%v'", db.TITLE_COLUMN, articleTitle, db.SECTION_COLUMN, a.Section),
	)
	if err != nil {
		slog.Error("Failed to find article", err)
		panic(err)
	}

	vars := ArticleVariables{
		Title:     article.GetString(db.TITLE_COLUMN),
		Author:    article.GetString(db.AUTHOR_COLUMN),
		Section:   article.GetString(db.SECTION_COLUMN),
		ImagePath: article.GetString(db.IMAGEPATH_COLUMN),
		Article:   article.GetString(db.ARTICLE_COLUMN),
		Created:   article.GetCreated().Time().Format(time.DateOnly),
		Updated:   article.GetUpdated().Time().Format(time.DateOnly),
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		slog.Error("Failed to execute template", err)
		panic(err)
	}
}
