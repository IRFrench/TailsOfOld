package weberrors

import (
	weberrors "TailsOfOld/TailsOfOld/handlers/web_errors"

	"github.com/go-chi/chi/v5"
)

func AddErrorRoutes(router *chi.Mux) {
	// Create error handler routes
	router.NotFound(weberrors.NotFoundHandler)
}
