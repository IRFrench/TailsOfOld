package main

import (
	"TailsOfOld/TailsOfOld/db"
	"TailsOfOld/TailsOfOld/server"
	"TailsOfOld/cfg"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

const (
	CONFIG_FILE = "config.yaml"
)

func main() {
	slog.Info("Building Server", "Version", "0.1.0")

	// Read configuration
	config, err := cfg.LoadConfig(CONFIG_FILE)
	if err != nil {
		slog.Error("Failed to read config", "Error", err)
		return
	}

	// Setup channels
	errorLog := make(chan error)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Make the database
	database := makeDatabase(config)

	// Create the server
	server, err := server.CreateServer(config, database)
	if err != nil {
		slog.Error("Server could not be created", "Error", err)
		return
	}

	// Run the DB
	go database.Start()

	// Run the server
	slog.Info("Running server", "Address", config.Web.Address)
	go server.Run(errorLog)

	// Listen to channels
	for {
		select {
		case err := <-errorLog:
			slog.Error("Server encountered an error", "Error", err)
			return
		case sig := <-sigs:
			slog.Info("Signal Recieved", "Signal", sig)
			if err := server.Shutdown(); err != nil {
				slog.Error("Error on Shutdown", "Error", err)
			}
			return
		}
	}
}

func makeDatabase(config cfg.Configuration) *pocketbase.PocketBase {
	app := pocketbase.NewWithConfig(
		pocketbase.Config{
			DefaultDataDir: config.Database.DataDir,
		},
	)
	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./database/pb_public"), false))
		return nil
	})

	app.OnModelAfterCreate(db.DB_ARTICLES).Add(func(e *core.ModelEvent) error {
		// Emails??
		return nil
	})

	return app
}
