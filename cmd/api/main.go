package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/evt/rest-api-example/repository"

	"github.com/evt/rest-api-example/mysqldb"
	echoLog "github.com/labstack/gommon/log"

	"github.com/evt/rest-api-example/gcloud"

	gcloudRepo "github.com/evt/rest-api-example/repository/gcloud"
	gcloudService "github.com/evt/rest-api-example/service/gcloud"

	"github.com/evt/rest-api-example/config"
	"github.com/evt/rest-api-example/controller"
	libError "github.com/evt/rest-api-example/lib/error"
	"github.com/evt/rest-api-example/lib/validator"
	"github.com/evt/rest-api-example/logger"
	"github.com/evt/rest-api-example/pgdb"

	mysqlrepo "github.com/evt/rest-api-example/repository/mysql"
	pgrepo "github.com/evt/rest-api-example/repository/pg"

	"github.com/evt/rest-api-example/service/web"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	// logger
	l := logger.Get()

	// connect to Postgres
	pgDB, err := pgdb.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations...")
		if err := runPgMigrations(cfg); err != nil {
			log.Fatal(err)
		}
	}

	// connect to MySQL
	mysqlDB, err := mysqldb.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run MySQL migrations
	if mysqlDB != nil {
		log.Println("Running MySQL migrations...")
		if err := runMysqlMigrations(cfg); err != nil {
			log.Fatal(err)
		}
	}

	// connect to google cloud
	cloudStorage, err := gcloud.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// store includes all repositories we have
	store := &repository.Store{}

	// Init Postgres repositories
	if pgDB != nil {
		store.User = pgrepo.NewUserRepo(pgDB)
		store.File = pgrepo.NewFileRepo(pgDB)
	}
	// Init MySQL repositories
	if mysqlDB != nil {
		store.User = mysqlrepo.NewUserRepo(mysqlDB)
		store.File = mysqlrepo.NewFileRepo(mysqlDB)
	}
	store.FileContent = gcloudRepo.NewFileRepo(cloudStorage, cfg.GCBucket)

	userService := web.NewUserWebService(ctx, store)
	fileService := web.NewFileWebService(ctx, store)
	fileContentService := gcloudService.NewFileContentService(ctx, store)

	// Init controllers
	userController := controller.NewUsers(ctx, userService, l)
	fileController := controller.NewFiles(ctx, fileService, fileContentService, l)

	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = libError.Error
	// Disable Echo JSON logger in debug mode
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API V1
	v1 := e.Group("/v1")

	// User routes
	userRoutes := v1.Group("/user")
	userRoutes.POST("/", userController.Create)
	userRoutes.GET("/:id", userController.Get)
	userRoutes.DELETE("/:id", userController.Delete)
	//userRoutes.PUT("/:id", userController.Update)

	// File routes
	fileRoutes := v1.Group("/file")
	fileRoutes.POST("/", fileController.Create)
	fileRoutes.GET("/:id", fileController.Get)
	fileRoutes.DELETE("/:id", fileController.Delete)
	fileRoutes.PUT("/:id/content", fileController.Upload)
	fileRoutes.GET("/:id/content", fileController.Download)

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}

// runPgMigrations runs Postgres migrations
func runPgMigrations(cfg *config.Config) error {
	if cfg.PgMigrationsPath == "" {
		return nil
	}
	if cfg.PgURL == "" {
		return errors.New("No cfg.PgURL provided")
	}
	m, err := migrate.New(
		cfg.PgMigrationsPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// runMysqlMigrations runs MySQL migrations
func runMysqlMigrations(cfg *config.Config) error {
	if cfg.MysqlMigrationsPath == "" {
		return nil
	}
	if cfg.MysqlDB == "" {
		return errors.New("No cfg.MysqlDB provided")
	}
	m, err := migrate.New(
		cfg.MysqlMigrationsPath,
		fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", cfg.MysqlUser, cfg.MysqlPassword, cfg.MysqlAddr, cfg.MysqlDB),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
