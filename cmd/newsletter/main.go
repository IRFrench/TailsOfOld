package main

import (
	"TailsOfOld/TailsOfOld/newsletter"
	"TailsOfOld/cfg"
	"os"
	"time"

	"github.com/pocketbase/pocketbase"
)

func main() {
	config, err := cfg.LoadConfig(os.Getenv("ETC"))
	if err != nil {
		panic(err)
	}

	mailClient := newsletter.NewMailClient(config)

	database := makeDatabase(config)
	go database.Start()

	time.Sleep(2 * time.Second)

	if err := mailClient.SendNewsletter(database, time.Now().AddDate(0, -1, 0)); err != nil {
		panic(err)
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
