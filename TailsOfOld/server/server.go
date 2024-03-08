package server

import (
	filesystem "TailsOfOld"
	"TailsOfOld/TailsOfOld/routes/articles"
	"TailsOfOld/TailsOfOld/routes/index"
	"context"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pocketbase/pocketbase"
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

func CreateServer(address string, database *pocketbase.PocketBase) (*WebServer, error) {
	router := chi.NewRouter()

	// Add static route
	staticFiles, err := filesystem.Static()
	if err != nil {
		return nil, err
	}

	fs.WalkDir(staticFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("failed")
		}
		fmt.Println(path)
		return nil
	})

	httpFileSystem := http.FS(staticFiles)
	staticHandler := http.FileServer(httpFileSystem)
	router.Handle("/static/*", staticHandler)

	// Add routes here
	index.AddIndexRoutes(router, database)
	articles.AddArticleOverviewRoutes(router, database)
	articles.AddArticleRoutes(router, database)

	// Add WebServer Deps 'ere
	return &WebServer{
		server: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}, nil
}
