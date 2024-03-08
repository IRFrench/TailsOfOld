package index

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

type IndexVariables struct {
	LatestGame        section.ArticleInfo
	LatestProgramming section.ArticleInfo
	LatestTale        section.ArticleInfo
}

type IndexHandler struct {
	Database *pocketbase.PocketBase
}

func (i IndexHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/index/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		slog.Error("Failed to parse file system into template", err)
		panic(err)
	}

	// Collect latest articles
	latestGamesArticle, err := i.Database.Dao().FindRecordsByFilter("articles", "section = 'games'", "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest games article from database", err)
		panic(err)
	}
	latestProgrammingArticle, err := i.Database.Dao().FindRecordsByFilter("articles", "section = 'programming'", "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest programming article from database", err)
		panic(err)
	}
	latestTalesArticle, err := i.Database.Dao().FindRecordsByFilter("articles", "section = 'tales'", "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest tales article from database", err)
		panic(err)
	}

	vars := IndexVariables{
		LatestProgramming: section.ArticleInfo{
			Title:       latestProgrammingArticle[0].GetString("title"),
			Date:        latestProgrammingArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestProgrammingArticle[0].GetString("author"),
			ImagePath:   latestProgrammingArticle[0].GetString("imagepath"),
			ArticlePath: fmt.Sprintf("/%v/%v", section.PROGRAMMING_SECTION, url.PathEscape(latestProgrammingArticle[0].GetString("title"))),
		},
		LatestGame: section.ArticleInfo{
			Title:       latestGamesArticle[0].GetString("title"),
			Date:        latestGamesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestGamesArticle[0].GetString("author"),
			ImagePath:   latestGamesArticle[0].GetString("imagepath"),
			ArticlePath: fmt.Sprintf("/%v/%v", section.GAME_SECTION, url.PathEscape(latestGamesArticle[0].GetString("title"))),
		},
		LatestTale: section.ArticleInfo{
			Title:       latestTalesArticle[0].GetString("title"),
			Date:        latestTalesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestTalesArticle[0].GetString("author"),
			ImagePath:   latestTalesArticle[0].GetString("imagepath"),
			ArticlePath: fmt.Sprintf("/%v/%v", section.TALES_SECTION, url.PathEscape(latestTalesArticle[0].GetString("title"))),
		},
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		slog.Error("Failed to execute template", err)
		panic(err)
	}
}
