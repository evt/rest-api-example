package store

import (
	"context"
	"log"

	"github.com/evt/rest-api-example/store/gcloud"

	"github.com/evt/rest-api-example/config"

	"github.com/pkg/errors"

	"github.com/evt/rest-api-example/store/mysql"
	"github.com/evt/rest-api-example/store/pg"
)

// Store contains all repositories
type Store struct {
	User        UserRepo
	File        FileRepo
	FileContent FileContentRepo
}

// New creates new store
func New(ctx context.Context, cfg *config.Config) (*Store, error) {
	// connect to google cloud
	cloudStorage, err := gcloud.Init(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "gcloud.Init failed")
	}

	// connect to Postgres
	pgDB, err := pg.Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	// Run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations...")
		if err := runPgMigrations(cfg); err != nil {
			return nil, errors.Wrap(err, "runPgMigrations failed")
		}
	}

	// connect to MySQL
	mysqlDB, err := mysql.Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "mysqldb.Dial failed")
	}

	// Run MySQL migrations
	if mysqlDB != nil {
		log.Println("Running MySQL migrations...")
		if err := runMysqlMigrations(cfg); err != nil {
			return nil, errors.Wrap(err, "runMysqlMigrations failed")
		}
	}

	var store Store

	// Init Postgres repositories
	if pgDB != nil {
		store.User = pg.NewUserRepo(pgDB)
		store.File = pg.NewFileRepo(pgDB)
	}
	// Init MySQL repositories
	if mysqlDB != nil {
		store.User = mysql.NewUserRepo(mysqlDB)
		store.File = mysql.NewFileRepo(mysqlDB)
	}
	store.FileContent = gcloud.NewFileRepo(cloudStorage, cfg.GCBucket)

	return &store, nil
}
