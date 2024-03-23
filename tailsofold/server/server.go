package server

import (
	filesystem "TailsOfOld"
	"TailsOfOld/cfg"
	"TailsOfOld/internal/db"
	"TailsOfOld/tailsofold/routes/articles"
	"TailsOfOld/tailsofold/routes/index"
	"TailsOfOld/tailsofold/routes/newsletter"
	"TailsOfOld/tailsofold/routes/search"
	weberrors "TailsOfOld/tailsofold/routes/web_errors"
	"context"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type WebServer struct {
	server *http.Server
}

func (s *WebServer) Run(errorChannel chan<- error) {
	if err := s.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			errorChannel <- err
		}
	}
}

func (s *WebServer) Shutdown() error {
	if err := s.server.Shutdown(context.Background()); err != nil {
		return err
	}
	return nil
}

func CreateServer(config cfg.Configuration, database *db.DatabaseClient) (*WebServer, error) {
	router := chi.NewRouter()
	newServer := &WebServer{
		server: &http.Server{
			Addr:    config.Web.Address,
			Handler: router,
		},
	}

	// Add static route
	staticFiles, err := filesystem.Static()
	if err != nil {
		return nil, err
	}

	fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Error().Err(err).Msg("failed to walk static files")
		}
		log.Debug().Msg(path)
		return nil
	})

	// Create static handler
	httpFileSystem := http.FS(staticFiles)
	staticHandler := http.FileServer(httpFileSystem)
	router.Handle("/static/*", staticHandler)

	// Create database handler
	databaseHandler := http.FileServer(http.Dir(config.Database.DataDir))
	router.Handle("/pb_data/*", http.StripPrefix("/pb_data", databaseHandler))

	// Add routes here
	if config.Web.Maintence {
		weberrors.AddMaintenceRoute(router)
		return newServer, nil
	}

	index.AddIndexRoutes(router, database)
	articles.AddArticleOverviewRoutes(router, database)
	articles.AddArticleRoutes(router, database)
	search.AddSearchRoutes(router, database)
	newsletter.AddNewsletterRoutes(router, database)
	weberrors.AddErrorRoutes(router)

	// Add WebServer Deps 'ere
	return newServer, nil
}
