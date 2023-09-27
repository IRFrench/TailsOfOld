package basic

import (
	"net/http"
	"text/template"
	"time"
)

type IndexVariables struct {
	Username string
	Date     string
}

func IndexHttp(response http.ResponseWriter, request *http.Request) {
	template := template.New("index.html")
	template, err := template.ParseFiles("TailsOfOld/templates/basic/index.html")
	if err != nil {
		panic(err)
	}

	vars := IndexVariables{Username: "Isaac", Date: time.Now().Format(time.DateTime)} //define an instance with required field

	if err := template.Execute(response, vars); err != nil {
		panic(err)
	}
}
