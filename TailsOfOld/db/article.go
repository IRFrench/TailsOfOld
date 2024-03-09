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
}
