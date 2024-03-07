package index

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/handlers"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/pocketbase/pocketbase"
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

type IndexHandler struct {
	Database *pocketbase.PocketBase
}

func (i IndexHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	templatePath := "TailsOfOld/static/templates/index/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(i.Database.Dao().FindRecordById("Articles", "j3e75qo8aejrjm8"))

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
