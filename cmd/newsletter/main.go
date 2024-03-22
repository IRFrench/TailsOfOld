package main

import (
	"TailsOfOld/cfg"
	"TailsOfOld/internal/db"
	"TailsOfOld/internal/newsletter"
	"os"
	"time"
)

func main() {
	config, err := cfg.LoadConfig(os.Getenv("ETC"))
	if err != nil {
		panic(err)
	}

	mailClient := newsletter.NewMailClient(config)

	database := db.NewDatabase(config)
	go database.Run()

	time.Sleep(2 * time.Second)

	if err := mailClient.SendNewsletter(database, time.Now().AddDate(0, -1, 0)); err != nil {
		panic(err)
	}
}
