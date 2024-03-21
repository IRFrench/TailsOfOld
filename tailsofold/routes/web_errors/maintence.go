package weberrors

import (
	weberrors "TailsOfOld/tailsofold/handlers/web_errors"

	"github.com/go-chi/chi/v5"
)

func AddMaintenceRoute(router *chi.Mux) {
	// Create index handler and route
	router.NotFound(weberrors.Maintence)
}
