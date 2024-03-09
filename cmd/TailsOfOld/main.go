package main

import (
	"TailsOfOld/TailsOfOld/server"
	"TailsOfOld/cfg"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pocketbase/pocketbase"
)

const (
	CONFIG_FILE = "config.yaml"
)

func main() {
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("version", "0.1.0").Msg("building server")

	// Read configuration
	config, err := cfg.LoadConfig(CONFIG_FILE)
	if err != nil {
		log.Error().Err(err).Msg("failed to read configuration")
		return
	}

	log.Info().
		Str("web address", config.Web.Address).
		Str("database dir", config.Database.DataDir).
		Msg("configuration loaded")

	// Setup channels
	errorLog := make(chan error)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Make the database
	database := makeDatabase(config)

	// Create the server
	server, err := server.CreateServer(config, database)
	if err != nil {
		log.Error().Err(err).Msg("failed to create server")
		return
	}

	// Run the DB
	go database.Start()

	// Run the server
	log.Info().Str("address", config.Web.Address)
	go server.Run(errorLog)

	// Listen to channels
	for {
		select {
		case err := <-errorLog:
			log.Error().Err(err).Msg("server encountered an error")
			return
		case sig := <-sigs:
			log.Info().Str("signal", sig.String()).Msg("signal recieved")
			if err := server.Shutdown(); err != nil {
				log.Error().Err(err).Msg("error on shutdown")
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

	return app
}
