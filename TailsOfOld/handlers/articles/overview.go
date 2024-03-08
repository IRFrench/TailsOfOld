package articles

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/handlers"
	"TailsOfOld/TailsOfOld/section"
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
	Articles []section.ArticleInfo
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
	databaseArticles, err := o.Database.Dao().FindRecordsByFilter("articles", fmt.Sprintf("section = '%v'", o.Section), "-created", 0, 0)
	if err != nil {
		slog.Error("Failed to get articles from database", err)
		panic(err)
	}

	allArticles := []section.ArticleInfo{}

	for _, article := range databaseArticles {
		allArticles = append(allArticles, section.ArticleInfo{
			Title:       article.GetString("title"),
			Description: article.GetString("description"),
			Date:        article.GetCreated().Time().Format(time.DateOnly),
			Author:      article.GetString("author"),
			ImagePath:   article.GetString("imagepath"),
			ArticlePath: fmt.Sprintf("/%v/%v", o.Section, url.PathEscape(article.GetString("title"))),
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
