package main

import (
	"context"

	"github.com/evt/rest-api-example/gcloud"

	gcloudRepo "github.com/evt/rest-api-example/repository/gcloud"
	gcloudService "github.com/evt/rest-api-example/service/gcloud"

	"github.com/evt/rest-api-example/config"
	"github.com/evt/rest-api-example/controller"
	"github.com/evt/rest-api-example/db"
	libError "github.com/evt/rest-api-example/lib/error"
	"github.com/evt/rest-api-example/lib/validator"
	"github.com/evt/rest-api-example/logger"
	"github.com/evt/rest-api-example/repository/pg"
	"github.com/evt/rest-api-example/service/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
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
	pgDB, err := db.Dial(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Run Postgres migrations
	log.Println("Running PostgreSQL migrations...")
	if err := runMigrations(cfg); err != nil {
		log.Fatal(err)
	}

	// connect to google cloud
	cloudStorage, err := gcloud.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Init repositories
	userRepo := pg.NewUserRepo(pgDB)
	fileRepo := pg.NewFileRepo(pgDB)
	fileContentRepo := gcloudRepo.NewFileRepo(cloudStorage, cfg.GCBucket)

	// Init services
	userService := web.NewUserWebService(ctx, userRepo)
	fileService := web.NewFileWebService(ctx, fileRepo)
	fileContentService := gcloudService.NewFileContentService(ctx, fileRepo, fileContentRepo)

	// Init controllers
	userController := controller.NewUsers(ctx, userService, l)
	fileController := controller.NewFiles(ctx, fileService, fileContentService, l)

	// Initialize Echo instance
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = libError.Error

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
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}

// runMigrations runs Postgres migrations
func runMigrations(cfg *config.Config) error {
	if cfg.PgMigrationsPath == "" {
		return errors.New("No cfg.PgMigrationsPath provided")
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
