package main

import (
	"TailsOfOld/TailsOfOld/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	ADDR = "127.0.0.1:9000"
)

func main() {
	slog.Info("Building Server", "Version", "0.1.0")

	errorLog := make(chan error)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := server.CreateServer(ADDR)

	slog.Info("Running server")
	go server.Run(errorLog)

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
