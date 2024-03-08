package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/handlers"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/pocketbase/pocketbase"
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
		slog.Error("Failed to parse file system into template", err)
		panic(err)
	}

	// Gather articles
	databaseArticles, err := o.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v'", db.SECTION_COLUMN, o.Section), "-created", 0, 0)
	if err != nil {
		slog.Error("Failed to get articles from database", err)
		panic(err)
	}

	allArticles := []db.ArticleInfo{}

	for _, article := range databaseArticles {
		allArticles = append(allArticles, db.ArticleInfo{
			Title:       article.GetString(db.TITLE_COLUMN),
			Description: article.GetString(db.DESCRIPTION_COLUMN),
			Date:        article.GetCreated().Time().Format(time.DateOnly),
			Author:      article.GetString(db.AUTHOR_COLUMN),
			ImagePath:   article.GetString(db.IMAGEPATH_COLUMN),
			ArticlePath: fmt.Sprintf("/%v/%v", o.Section, url.PathEscape(article.GetString(db.TITLE_COLUMN))),
		})
	}

	vars := OverviewVariables{
		Section:  o.Section,
		Articles: allArticles,
	}

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		slog.Error("Failed to execute template", err)
		panic(err)
	}
}
