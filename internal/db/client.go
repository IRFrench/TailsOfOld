package db

import (
	"TailsOfOld/cfg"

	"github.com/pocketbase/pocketbase"
)

type DatabaseClient struct {
	Db *pocketbase.PocketBase
}

func NewDatabase(config cfg.Configuration) *DatabaseClient {
	app := pocketbase.NewWithConfig(
		pocketbase.Config{
			DefaultDataDir: config.Database.DataDir,
		},
	)
	return &DatabaseClient{
		Db: app,
	}
}

func (d *DatabaseClient) Run() error {
	return d.Db.Start()
}
