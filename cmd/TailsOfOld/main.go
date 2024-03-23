package main

import (
	"TailsOfOld/cfg"
	"TailsOfOld/internal/db"
	mailclient "TailsOfOld/internal/mail_client"
	"TailsOfOld/tailsofold/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("version", "0.1.0").Msg("building server")

	// Read configuration
	config, err := cfg.LoadConfig(os.Getenv("ETC"))
	if err != nil {
		log.Err(err).Msg("failed to read configuration")
		return
	}

	// Setup channels
	errorLog := make(chan error)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Make the database
	database := db.NewDatabase(config)

	// Created the Mail client
	mail := mailclient.NewMailClient(config)

	// Create the server
	server, err := server.CreateServer(config, database, mail)
	if err != nil {
		log.Err(err).Msg("failed to create server")
		return
	}

	// Run the DB
	go database.Run()

	// Run the server
	log.Info().Str("address", config.Web.Address)
	go server.Run(errorLog)

	// Listen to channels
	for {
		select {
		case err := <-errorLog:
			log.Err(err).Msg("server encountered an error")
			return
		case sig := <-sigs:
			log.Info().Str("signal", sig.String()).Msg("signal recieved")
			if err := server.Shutdown(); err != nil {
				log.Err(err).Msg("error on shutdown")
			}
			return
		}
	}
}
