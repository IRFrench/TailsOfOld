package basic

import (
	"TailsOfOld/TailsOfOld/handlers/basic"

	"github.com/go-chi/chi/v5"
)

func AddIndexRoutes(router *chi.Mux) {
	router.HandleFunc("/", basic.IndexHttp)
}
