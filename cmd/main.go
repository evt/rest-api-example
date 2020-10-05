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

	//// router
	//router := mux.NewRouter()
	//
	//
	//// Repositories
	//pscRepo := pg.NewPSCPgRepo(pgDB)
	//guardRepo := pg.NewGuardPgRepo(pgDB)
	//crewRepo := pg.NewCrewPgRepo(pgDB)
	//userRepo := pg.NewUserPgRepo(pgDB)
	//orderRepo := pg.NewOrderPgRepo(pgDB)
	//orderCommentRepo := pg.NewOrderCommentPgRepo(pgDB)
	//orderFileRepo := pg.NewOrderFilePgRepo(pgDB)
	//tokenRepo := pg.NewTokenPgRepo(pgDB)
	//sessionRepo := pg.NewSessionPgRepo(pgDB)
	//fileMetadataRepo := pg.NewFileMetadataPgRepo(pgDB)
	//userPhotoRepo := pg.NewUserPhotoPgRepo(pgDB)
	//userSMSRepo := pg.NewUserSMSPgRepo(pgDB)
	//contactRepo := pg.NewContactPgRepo(pgDB)
	//contractRepo := pg.NewContractPgRepo(pgDB)
	//fileCloudRepo := yos.NewCloudStorageRepo(cloudStorage)
	//
	//// Services
	//pscService := service.NewPSCSvc(ctx, pscRepo, guardRepo, userRepo, contractRepo)
	//guardService := service.NewGuardSvc(ctx, guardRepo, pscRepo)
	//crewService := service.NewCrewSvc(ctx, crewRepo, pscRepo, guardRepo)
	//userService := service.NewUserSvc(ctx, userRepo, sessionRepo, tokenRepo, fileMetadataRepo, userPhotoRepo, userSMSRepo)
	//orderService := service.NewOrderSvc(ctx, orderRepo, crewRepo, guardRepo, pscRepo, userRepo)
	//orderCommentService := service.NewOrderCommentSvc(ctx, orderCommentRepo, orderRepo, userRepo)
	//orderFileService := service.NewOrderFileSvc(ctx, orderFileRepo, orderRepo, userRepo, fileMetadataRepo)
	//tokenService := service.NewTokenSvc(ctx, tokenRepo, pscRepo, userRepo)
	//sessionService := service.NewSessionSvc(ctx, sessionRepo, tokenRepo, userRepo)
	//fileMetadataService := service.NewFileMetadataSvc(ctx, fileMetadataRepo, fileCloudRepo)
	//userPhotoService := service.NewUserPhotoSvc(ctx, userPhotoRepo, userRepo)
	//userSMSService := service.NewUserSMSSvc(ctx, userSMSRepo, userRepo)
	//contactService := service.NewContactSvc(ctx, contactRepo, userRepo)
	//contractService := service.NewContractSvc(ctx, contractRepo, userRepo, pscRepo, fileMetadataRepo, fileCloudRepo)
	//geoService := service.NewGeoSvc(ctx)
	//fileService := service.NewFileSvc(ctx, fileMetadataRepo, fileCloudRepo)
	//
	//// Set up server
	//server := &server.Server{
	//	Config:              cfg,
	//	PSCService:          pscService,
	//	GuardService:        guardService,
	//	CrewService:         crewService,
	//	UserService:         userService,
	//	OrderService:        orderService,
	//	OrderCommentService: orderCommentService,
	//	OrderFileService:    orderFileService,
	//	TokenService:        tokenService,
	//	SessionService:      sessionService,
	//	FileMetadataService: fileMetadataService,
	//	UserPhotoService:    userPhotoService,
	//	UserSMSService:      userSMSService,
	//	ContactService:      contactService,
	//	ContractService:     contractService,
	//	GeoService:          geoService,
	//	FileService:         fileService,
	//	Router:              router,
	//	Logger:              zlog,
	//}
	//
	//// Run http server
	//server.Routes()
	//httpServer := &http.Server{
	//	Addr:    cfg.HTTPAddr,
	//	Handler: server,
	//}
	//
	//stopHTTPServer := make(chan bool)
	//
	//log.Printf("GuardPoint server (HTTP) listening on %s", cfg.HTTPAddr)
	//if err := httpServer.ListenAndServe(); err != nil {
	//	zlog.Sugar().Fatal(err)
	//	stopHTTPServer <- true
	//}
	//
	//<-stopHTTPServer

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
