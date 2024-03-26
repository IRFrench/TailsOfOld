package db

import (
	"TailsOfOld/cfg"

	"github.com/pocketbase/pocketbase"
)

type DatabaseClient struct {
	Db *pocketbase.PocketBase
}

func NewDatabase(config cfg.Database) *DatabaseClient {
	app := pocketbase.NewWithConfig(
		pocketbase.Config{
			DefaultDataDir: config.DataDir,
		},
	)
	return &DatabaseClient{
		Db: app,
	}
}

func (d *DatabaseClient) Run(errorChannel chan<- error) {
	if err := d.Db.Start(); err != nil {
		errorChannel <- err
	}
}
