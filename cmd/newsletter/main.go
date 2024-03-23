package main

import (
	"TailsOfOld/cfg"
	"TailsOfOld/internal/db"
	mailclient "TailsOfOld/internal/mail_client"
	"os"
	"time"
)

func main() {
	config, err := cfg.LoadConfig(os.Getenv("ETC"))
	if err != nil {
		panic(err)
	}

	mailClient := mailclient.NewMailClient(config)

	database := db.NewDatabase(config)
	go database.Run()

	time.Sleep(2 * time.Second)

	if err := mailClient.SendNewsletter(database, time.Now().AddDate(0, -1, 0)); err != nil {
		panic(err)
	}
}
