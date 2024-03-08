package section

const (
	GAME_SECTION        = "games"
	PROGRAMMING_SECTION = "programming"
	TALES_SECTION       = "tales"
)

type ArticleInfo struct {
	Title       string
	Description string
	Date        string
	Author      string
	ImagePath   string
	ArticlePath string
}
