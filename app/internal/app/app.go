package app

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"javacode-test-task/app/internal/config"
	"javacode-test-task/app/internal/wallet"
	"javacode-test-task/app/internal/wallet/db"
	"javacode-test-task/app/pkg/client/psql"
	"javacode-test-task/app/pkg/middlewares"
	"log/slog"
	"net/http"
)

type App struct {
	config     *config.Config
	logger     *slog.Logger
	router     *httprouter.Router
	httpServer *http.Server
}

func NewApp(cfg *config.Config, log *slog.Logger) (*App, error) {
	log.Info("Try to initialize app")
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Storage.DBUser,
		cfg.Storage.DBPassword,
		cfg.Storage.DBHost,
		cfg.Storage.DBPort,
		cfg.Storage.DBName,
	)
	log.Info("Try to connect to database")
	psqlClient, err := psql.NewClient(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	log.Info("Success connect to database")
	storage := db.NewRepository(psqlClient, log)
	s := wallet.NewService(storage)

	v := validator.New()
	router := httprouter.New()
	router.GET("/api/v1/wallets/", middlewares.LoggingMiddleware(log, wallet.GetAllWallets(s)))
	router.GET("/api/v1/wallets/:id", middlewares.LoggingMiddleware(log, wallet.FindWalletByUUID(s)))
	router.POST("/api/v1/wallet", middlewares.LoggingMiddleware(log, wallet.UpdateWalletByUUID(v, s)))
	app := App{
		config: cfg,
		logger: log,
		router: router,
	}
	log.Info("Success initialize app")
	return &app, nil
}

func (app *App) StartHTTP() {
	app.logger.Info("Start HTTP server")
	addr := fmt.Sprintf("%s:%s", app.config.Server.Address, app.config.Server.Port)
	app.httpServer = &http.Server{
		Addr:    addr,
		Handler: app.router,
	}

	if err := app.httpServer.ListenAndServe(); err != nil {
		app.logger.Error("Failed to start HTTP server", err)
	}
}
