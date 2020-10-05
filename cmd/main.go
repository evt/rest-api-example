package main

import (
	"context"
	"github.com/evt/simple-web-server/config"
	"github.com/evt/simple-web-server/controller"
	"github.com/evt/simple-web-server/db"
	"github.com/evt/simple-web-server/logger"
	"github.com/evt/simple-web-server/repository/pg"
	"github.com/evt/simple-web-server/service/web"
	"github.com/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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
	_ = logger.Get()

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

	// Init repositories
	userRepo := pg.NewUserRepo(pgDB)

	// Init services
	userService := web.NewUserWebService(ctx, userRepo)

	// Init controllers
	userController := controller.NewUsers(ctx, userService)

	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	userRoutes := e.Group("/users")
	userRoutes.GET("/", userController.Get)

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

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
