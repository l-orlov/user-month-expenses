package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/l-orlov/user-month-expenses/internal/config"
	"github.com/l-orlov/user-month-expenses/internal/handler"
	"github.com/l-orlov/user-month-expenses/internal/repository"
	"github.com/l-orlov/user-month-expenses/internal/repository/postgres"
	"github.com/l-orlov/user-month-expenses/internal/server"
	"github.com/l-orlov/user-month-expenses/internal/service"
	"github.com/l-orlov/user-month-expenses/pkg/logger"
	_ "github.com/lib/pq"
)

// Run initializes whole application.
func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	lg, err := logger.New(cfg.Logger.Level, cfg.Logger.Format)
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}

	// Dependencies
	db, err := postgres.ConnectToDB(cfg.PostgresDB)
	if err != nil {
		lg.Fatalf("failed to connect to db: %v", err)
	}
	defer func() {
		if err = db.Close(); err != nil {
			lg.Errorf("failed to close db: %v", err)
		}
	}()

	if cfg.PostgresDB.MigrationMode {
		if err = postgres.MigrateSchema(db.DB, cfg.PostgresDB); err != nil {
			log.Fatalf("failed to do migration: %v", err)
		}
	}

	// Repo, Service & API Handlers
	repo := repository.NewRepository(cfg, db)
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	svc := service.NewService(repo)
	if err != nil {
		log.Fatalf("failed to create service: %v", err)
	}

	h := handler.New(cfg, lg, svc)

	// HTTP Server
	srv := server.New(cfg.Port, h.InitRoutes())
	go func() {
		if err = srv.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			lg.Fatalf("error occurred while running http server: %v", err)
		}
	}()

	lg.Infof("service started on port %s", cfg.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	lg.Info("service shutting down")

	if err = srv.Shutdown(context.Background()); err != nil {
		lg.Errorf("failed to shut down: %v", err)
	}
}
