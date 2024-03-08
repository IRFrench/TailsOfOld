package db

const (
	DB_ARTICLES = "articles"

	GAME_SECTION        = "games"
	PROGRAMMING_SECTION = "programming"
	TALES_SECTION       = "tales"

	TITLE_COLUMN       = "title"
	DESCRIPTION_COLUMN = "description"
	AUTHOR_COLUMN      = "author"
	SECTION_COLUMN     = "section"
	IMAGEPATH_COLUMN   = "imagepath"
	ARTICLE_COLUMN     = "article"
)

type ArticleInfo struct {
	Title       string
	Description string
	Date        string
	Author      string
	ImagePath   string
	ArticlePath string
}
