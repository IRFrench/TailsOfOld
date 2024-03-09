package index

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

type IndexVariables struct {
	LatestGame        db.ArticleInfo
	LatestProgramming db.ArticleInfo
	LatestTale        db.ArticleInfo
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
	latestGamesArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.GAME_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest games article from database", err)
		panic(err)
	}
	latestProgrammingArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.PROGRAMMING_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest programming article from database", err)
		panic(err)
	}
	latestTalesArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.TALES_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		slog.Error("Failed to gather latest tales article from database", err)
		panic(err)
	}

	vars := IndexVariables{
		LatestProgramming: db.ArticleInfo{
			Title:       latestProgrammingArticle[0].GetString(db.TITLE_COLUMN),
			Date:        latestProgrammingArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestProgrammingArticle[0].GetString(db.AUTHOR_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestProgrammingArticle[0].BaseFilesPath(), latestProgrammingArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.PROGRAMMING_SECTION, url.PathEscape(latestProgrammingArticle[0].GetString("title"))),
		},
		LatestGame: db.ArticleInfo{
			Title:       latestGamesArticle[0].GetString(db.TITLE_COLUMN),
			Date:        latestGamesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestGamesArticle[0].GetString(db.AUTHOR_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestGamesArticle[0].BaseFilesPath(), latestGamesArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.GAME_SECTION, url.PathEscape(latestGamesArticle[0].GetString("title"))),
		},
		LatestTale: db.ArticleInfo{
			Title:       latestTalesArticle[0].GetString(db.TITLE_COLUMN),
			Date:        latestTalesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      latestTalesArticle[0].GetString(db.AUTHOR_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestTalesArticle[0].BaseFilesPath(), latestTalesArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.TALES_SECTION, url.PathEscape(latestTalesArticle[0].GetString("title"))),
		},
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		slog.Error("Failed to execute template", err)
		panic(err)
	}
}
