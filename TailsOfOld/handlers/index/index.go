package index

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
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		weberrors.Borked(response, request)
		return
	}

	// Collect latest games article
	latestGamesArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.GAME_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest games article from database")
		weberrors.Borked(response, request)
		return
	}
	gamesAuthor, err := i.Database.Dao().FindRecordById(db.DB_USERS, latestGamesArticle[0].GetString(db.AUTHOR_COLUMN))
	if err != nil {
		log.Error().Err(err).Str("game article title", latestGamesArticle[0].GetString(db.TITLE_COLUMN)).Msg("failed to find game author")
		weberrors.Borked(response, request)
		return
	}

	// Collect latest programming article
	latestProgrammingArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.PROGRAMMING_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest programming article from database")
		weberrors.Borked(response, request)
		return
	}
	programmingAuthor, err := i.Database.Dao().FindRecordById(db.DB_USERS, latestProgrammingArticle[0].GetString(db.AUTHOR_COLUMN))
	if err != nil {
		log.Error().Err(err).Str("programming article title", latestProgrammingArticle[0].GetString(db.TITLE_COLUMN)).Msg("failed to find programming author")
		weberrors.Borked(response, request)
		return
	}

	// Collect latest tales article
	latestTalesArticle, err := i.Database.Dao().FindRecordsByFilter(db.DB_ARTICLES, fmt.Sprintf("%v = '%v' && %v = true", db.SECTION_COLUMN, db.TALES_SECTION, db.LIVE_COLUMN), "-created", 1, 0)
	if err != nil {
		log.Error().Err(err).Msg("failed to gather latest tales article from database")
		weberrors.Borked(response, request)
		return
	}
	talesAuthor, err := i.Database.Dao().FindRecordById(db.DB_USERS, latestTalesArticle[0].GetString(db.AUTHOR_COLUMN))
	if err != nil {
		log.Error().Err(err).Str("tale article title", latestTalesArticle[0].GetString(db.TITLE_COLUMN)).Msg("failed to find tale author")
		weberrors.Borked(response, request)
		return
	}

	vars := IndexVariables{
		LatestProgramming: db.ArticleInfo{
			Title:       latestProgrammingArticle[0].GetString(db.TITLE_COLUMN),
			Created:     latestProgrammingArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      programmingAuthor.GetString(db.NAME_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestProgrammingArticle[0].BaseFilesPath(), latestProgrammingArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.PROGRAMMING_SECTION, url.PathEscape(latestProgrammingArticle[0].GetString("title"))),
		},
		LatestGame: db.ArticleInfo{
			Title:       latestGamesArticle[0].GetString(db.TITLE_COLUMN),
			Created:     latestGamesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      gamesAuthor.GetString(db.NAME_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestGamesArticle[0].BaseFilesPath(), latestGamesArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.GAME_SECTION, url.PathEscape(latestGamesArticle[0].GetString("title"))),
		},
		LatestTale: db.ArticleInfo{
			Title:       latestTalesArticle[0].GetString(db.TITLE_COLUMN),
			Created:     latestTalesArticle[0].GetCreated().Time().Format(time.DateOnly),
			Author:      talesAuthor.GetString(db.NAME_COLUMN),
			ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", latestTalesArticle[0].BaseFilesPath(), latestTalesArticle[0].GetString(db.IMAGEPATH_COLUMN)),
			ArticlePath: fmt.Sprintf("/%v/%v", db.TALES_SECTION, url.PathEscape(latestTalesArticle[0].GetString("title"))),
		},
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		weberrors.Borked(response, request)
		return
	}
}
