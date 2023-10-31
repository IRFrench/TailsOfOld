package basic

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/handlers"
	"html/template"
	"net/http"
	"time"
)

type IndexVariables struct {
	Username string
	Date     string
}

func IndexHttp(response http.ResponseWriter, request *http.Request) {
	templatePath := "TailsOfOld/static/templates/basic/index.html"
	template := template.New("index.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		panic(err)
	}

	vars := IndexVariables{Username: "Isaac", Date: time.Now().Format(time.DateTime)} //define an instance with required field

	if err := template.ExecuteTemplate(response, "base", vars); err != nil {
		panic(err)
	}
}
