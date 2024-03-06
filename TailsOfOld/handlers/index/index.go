package index

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/handlers"
	"html/template"
	"net/http"
	"time"
)

type IndexVariables struct {
	LatestGame        ArticleInfo
	LatestProgramming ArticleInfo
	LatestTale        ArticleInfo
}

type ArticleInfo struct {
	Title       string
	Date        string
	Author      string
	ImagePath   string
	ArticlePath string
}

func IndexHttp(response http.ResponseWriter, request *http.Request) {
	templatePath := "TailsOfOld/static/templates/index/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		panic(err)
	}

	vars := IndexVariables{
		LatestProgramming: ArticleInfo{
			Title:       "Test Article Title",
			Date:        time.Now().Format(time.DateOnly),
			Author:      "Isaac French",
			ImagePath:   "/static/img/logo.png",
			ArticlePath: "#",
		},
		LatestGame: ArticleInfo{
			Title:       "Test Game Title",
			Date:        time.Now().Format(time.DateOnly),
			Author:      "Isaac French",
			ImagePath:   "https://image.api.playstation.com/vulcan/ap/rnd/202309/0718/ca77865b4bc8a1ea110fbe1492f7de8f80234dd079fc181a.png",
			ArticlePath: "#",
		},
		LatestTale: ArticleInfo{
			Title:       "Test Story Title",
			Date:        time.Now().Format(time.DateOnly),
			Author:      "Isaac French",
			ImagePath:   "/static/img/book_pup.jpeg",
			ArticlePath: "#",
		},
	} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		panic(err)
	}
}
