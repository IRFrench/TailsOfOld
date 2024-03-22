package weberrors

import (
	filesystem "TailsOfOld"
	"TailsOfOld/tailsofold/handlers"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

func NotFoundHandler(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "tailsofold/static/templates/error/404.html"
	template := template.New("404.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		Borked(response, request)
		return
	}

	if err := template.ExecuteTemplate(response, "base", nil); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		Borked(response, request)
		return
	}
}
