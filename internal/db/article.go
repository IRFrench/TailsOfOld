package db

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

const (
	DB_ARTICLES = "articles"

	GAME_SECTION        = "games"
	PROGRAMMING_SECTION = "programming"
	TALES_SECTION       = "tales"

	TITLE_COLUMN       = "title"
	DESCRIPTION_COLUMN = "description"
	AUTHOR_COLUMN      = "author"
	SECTION_COLUMN     = "section"
	IMAGEPATH_COLUMN   = "image"
	ARTICLE_COLUMN     = "article"
	LIVE_COLUMN        = "live"
)

type ArticleInfo struct {
	Title       string
	Section     string
	Description string
	Created     string
	Updated     string
	Author      string
	ImagePath   string
	ArticlePath string
	Article     string
	New         bool
}

func (d *DatabaseClient) GetLatestSectionArticleInfo(section string) (ArticleInfo, error) {
	latestSectionArticle, err := d.Db.Dao().FindRecordsByFilter(
		DB_ARTICLES,
		fmt.Sprintf("%v = '%v' && %v = true", SECTION_COLUMN, section, LIVE_COLUMN),
		"-created",
		1,
		0,
	)
	if err != nil {
		return ArticleInfo{}, err
	}

	return parseArticle(latestSectionArticle[0]), nil
}

func (d *DatabaseClient) GetEntireSectionArticleInfo(section string) ([]ArticleInfo, error) {
	sectionArticles, err := d.Db.Dao().FindRecordsByFilter(
		DB_ARTICLES,
		fmt.Sprintf("%v = '%v' && %v = true", SECTION_COLUMN, section, LIVE_COLUMN),
		"-created",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	allArticles := make([]ArticleInfo, len(sectionArticles))
	for index := range sectionArticles {
		allArticles[index] = parseArticle(sectionArticles[index])
	}

	return allArticles, nil
}

func (d *DatabaseClient) SearchArticlesByTitle(title string) ([]ArticleInfo, error) {
	articles, err := d.Db.Dao().FindRecordsByExpr(
		DB_ARTICLES,
		dbx.Like(TITLE_COLUMN, title),
		dbx.NewExp(fmt.Sprintf("%v = true", LIVE_COLUMN)),
	)
	if err != nil {
		return []ArticleInfo{}, nil
	}

	allArticles := make([]ArticleInfo, len(articles))
	for index := range articles {
		allArticles[index] = parseArticle(articles[index])
	}

	return allArticles, nil
}

func (d *DatabaseClient) GetFullArticle(title, section string) (ArticleInfo, error) {
	article, err := d.Db.Dao().FindFirstRecordByFilter(
		DB_ARTICLES,
		fmt.Sprintf("%v = {:title} && %v = '%v' && %v = true", TITLE_COLUMN, SECTION_COLUMN, section, LIVE_COLUMN),
		dbx.Params{"title": title},
	)
	if err != nil {
		return ArticleInfo{}, err
	}

	fullArticle := parseArticle(article)
	fullArticle.Article = article.GetString(ARTICLE_COLUMN)

	return fullArticle, nil
}

func (d *DatabaseClient) GetArticlesCreatedSinceTime(since time.Time) ([]ArticleInfo, error) {
	articles, err := d.Db.Dao().FindRecordsByFilter(
		DB_ARTICLES,
		fmt.Sprintf("created >= '%v' && %v = true", since.Format(time.DateTime), LIVE_COLUMN),
		"-created",
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	allArticles := make([]ArticleInfo, len(articles))
	for index := range articles {
		allArticles[index] = parseArticle(articles[index])
	}

	return allArticles, nil
}

func parseArticle(article *models.Record) ArticleInfo {
	return ArticleInfo{
		Title:       article.GetString(TITLE_COLUMN),
		Section:     article.GetString(SECTION_COLUMN),
		Description: article.GetString(DESCRIPTION_COLUMN),
		Created:     article.GetCreated().Time().Format(time.DateOnly),
		Updated:     article.GetUpdated().Time().Format(time.DateOnly),
		Author:      article.GetString(AUTHOR_COLUMN),
		ImagePath:   fmt.Sprintf("/pb_data/storage/%v/%v", article.BaseFilesPath(), article.GetString(IMAGEPATH_COLUMN)),
		ArticlePath: fmt.Sprintf("/%v/%v", article.GetString(SECTION_COLUMN), url.PathEscape(article.GetString(TITLE_COLUMN))),
		New:         time.Now().Before(article.GetCreated().Time().AddDate(0, 1, 0)),
	}
}
