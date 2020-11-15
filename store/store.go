package store

import (
	"context"
	"log"
	"time"

	"github.com/evt/rest-api-example/store/local"

	"github.com/evt/rest-api-example/logger"

	"github.com/evt/rest-api-example/store/gcloud"

	"github.com/evt/rest-api-example/config"

	"github.com/pkg/errors"

	"github.com/evt/rest-api-example/store/mysql"
	"github.com/evt/rest-api-example/store/pg"
)

// Store contains all repositories
type Store struct {
	Pg    *pg.DB       // for KeepAlivePg (see below)
	MySQL *mysql.MySQL // for KeepAliveMySQL (see below)

	User        UserRepo
	File        FileMetaRepo
	FileContent FileContentRepo
}

// New creates new store
func New(ctx context.Context) (*Store, error) {
	cfg := config.Get()

	// connect to Postgres
	pgDB, err := pg.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	// Run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations...")
		if err := runPgMigrations(); err != nil {
			return nil, errors.Wrap(err, "runPgMigrations failed")
		}
	}

	// connect to MySQL
	mysqlDB, err := mysql.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "mysqldb.Dial failed")
	}

	// Run MySQL migrations
	if mysqlDB != nil {
		log.Println("Running MySQL migrations...")
		if err := runMysqlMigrations(); err != nil {
			return nil, errors.Wrap(err, "runMysqlMigrations failed")
		}
	}

	var store Store

	// Init Postgres repositories
	if pgDB != nil {
		store.Pg = pgDB
		go store.KeepAlivePg()
		store.User = pg.NewUserRepo(pgDB)
		store.File = pg.NewFileMetaRepo(pgDB)
	}
	// Init MySQL repositories
	if mysqlDB != nil {
		store.MySQL = mysqlDB
		go store.KeepAliveMySQL()
		store.User = mysql.NewUserRepo(mysqlDB)
		store.File = mysql.NewFileMetaRepo(mysqlDB)
	}

	switch {
	case cfg.FilePath != "":
		store.FileContent = local.NewFileContentRepo(cfg.FilePath)

	// connect to google cloud if bucket defined in config
	case cfg.GCBucket != "":
		cloudStorage, err := gcloud.Init(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "gcloud.Init failed")
		}
		store.FileContent = gcloud.NewFileContentRepo(cloudStorage, cfg.GCBucket)
	}

	return &store, nil
}

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func (store *Store) KeepAlivePg() {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.Pg == nil {
			lostConnect = true
		} else if _, err = store.Pg.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		store.Pg, err = pg.Dial()
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAlivePg] PostgreSQL reconnected")
	}
}

// KeepAliveMySQL makes sure MySQL is alive and reconnects if needed
func (store *Store) KeepAliveMySQL() {
	logger := logger.Get()
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if store.MySQL == nil {
			lostConnect = true
		} else if err = store.MySQL.DB.DB().Ping(); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Debug().Msg("[store.KeepAliveMySQL] Lost MySQL connection. Restoring...")
		store.MySQL, err = mysql.Dial()
		if err != nil {
			logger.Err(err)
			continue
		}
		logger.Debug().Msg("[store.KeepAliveMySQL] MySQL reconnected")
	}
}
