package weberrors

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/handlers"
	"net/http"
	"text/template"

	"github.com/rs/zerolog/log"
)

func Maintence(response http.ResponseWriter, request *http.Request) {
	// Build template
	templatePath := "TailsOfOld/static/templates/error/maintence.html"
	template := template.New("maintence.html")

	template, err := template.ParseFS(filesystem.FileSystem, handlers.BASE_TEMPLATES, templatePath)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse filesystem into the template")
		panic(err)
	}

	if err := template.ExecuteTemplate(response, "base", nil); err != nil {
		log.Error().Err(err).Msg("failed to execute the template")
		panic(err)
	}
}