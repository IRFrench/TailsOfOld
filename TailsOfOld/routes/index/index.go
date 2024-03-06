package index

import (
	"TailsOfOld/TailsOfOld/handlers/index"

	"github.com/go-chi/chi/v5"
)

func AddIndexRoutes(router *chi.Mux) {
	router.HandleFunc("/", index.IndexHttp)
}
