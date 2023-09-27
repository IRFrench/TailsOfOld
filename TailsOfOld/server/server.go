package server

import (
	"TailsOfOld/TailsOfOld/routes/basic"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func CreateServer(address string) *WebServer {
	router := chi.NewRouter()

	// Add routes here
	basic.AddIndexRoutes(router)

	// Add WebServer Deps 'ere

	return &WebServer{
		server: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}
}
