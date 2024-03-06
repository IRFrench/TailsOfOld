package articles

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type sections string

const (
	GAMES_SECTION       = sections("games")
	PROGRAMMING_SECTION = sections("programming")
	TALES_SECTION       = sections("tales")
)

var (
	ErrFailedToReadArticlesFile = errors.New("failed to read articles file")
	ErrFailedToReadFileContent  = errors.New("failed to read articles file content")
	ErrFailedToWriteJson        = errors.New("failed to write articles in json")
	ErrFailedToUpdateFile       = errors.New("failed to write articles json file")
	ErrArticleExists            = errors.New("article already exists")
)

type Articles struct {
	Programming []ArticleMetaData `json:"programming"`
	Games       []ArticleMetaData `json:"games"`
	Tales       []ArticleMetaData `json:"tales"`
}

type ArticleMetaData struct {
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Author      string    `json:"author"`
	ArticlePath string    `json:"path"`
	ImagePath   string    `json:"image"`
}

type ArticleTracker struct {
	ArticleFilePath   string
	ArticleFolderPath string
	Articles          Articles
}

func (a *ArticleTracker) ReadArticlesFile() error {
	fileContent, err := os.ReadFile(a.ArticleFilePath)
	if err != nil {
		return ErrFailedToReadArticlesFile
	}

	if err := json.Unmarshal(fileContent, &a.Articles); err != nil {
		return ErrFailedToReadFileContent
	}

	return nil
}

func (a *ArticleTracker) CreateArticle(date time.Time, title, author, image string, section sections) error {
	newArticle := ArticleMetaData{
		Title:       title,
		Date:        date,
		Author:      author,
		ImagePath:   image,
		ArticlePath: fmt.Sprintf("%v/%v/%v.html", a.ArticleFolderPath, string(section), title),
	}

	for _, article := range a.GetSectionArticles(section) {
		if newArticle.Title == article.Title {
			return ErrArticleExists
		}
	}

	if section == GAMES_SECTION {
		a.Articles.Games = append(a.Articles.Games, newArticle)
	}
	if section == PROGRAMMING_SECTION {
		a.Articles.Programming = append(a.Articles.Programming, newArticle)
	}
	if section == TALES_SECTION {
		a.Articles.Tales = append(a.Articles.Tales, newArticle)
	}

	articlesJson, err := json.Marshal(a.Articles)
	if err != nil {
		return ErrFailedToWriteJson
	}

	if err = os.WriteFile(a.ArticleFilePath, articlesJson, os.ModePerm); err != nil {
		return ErrFailedToUpdateFile
	}

	return os.WriteFile(newArticle.ArticlePath, nil, os.ModePerm)
}

func (a *ArticleTracker) GetSectionArticles(section sections) []ArticleMetaData {
	return a.getSection(section)
}

func (a *ArticleTracker) GetLatestArticle(section sections) ArticleMetaData {
	sectionArticles := a.getSection(section)
	latestArticle := ArticleMetaData{}

	for _, article := range sectionArticles {
		if article.Date.Before(latestArticle.Date) {
			continue
		}
		latestArticle = article
	}

	return latestArticle
}

func (a *ArticleTracker) getSection(section sections) []ArticleMetaData {
	if section == GAMES_SECTION {
		return a.Articles.Games
	}
	if section == PROGRAMMING_SECTION {
		return a.Articles.Programming
	}
	if section == TALES_SECTION {
		return a.Articles.Tales
	}
	return nil
}

func NewArticleTracker(filePath string, folderPath string) ArticleTracker {
	return ArticleTracker{
		ArticleFilePath:   filePath,
		ArticleFolderPath: folderPath,
	}
}
