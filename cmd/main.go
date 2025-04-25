package main

import (
	"context"
	"log/slog"
	"os"
	"test-people/internal/config"
	"test-people/internal/db"
	"test-people/internal/server"
)

// @title           People API
// @version         1.0
// @description     This is a sample server for managing people.
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.Load()

	slog.Info("Connecting to database...")
	conn, err := db.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error("Failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		if err := conn.Close(context.Background()); err != nil {
			slog.Warn("Failed to close db connection", slog.String("error", err.Error()))
		}
	}()

	slog.Info("Running database migrations...")
	if err := db.RunMigrations(cfg.Postgres, "/app/migrations"); err != nil {
		slog.Error("Migration error", slog.String("error", err.Error()))
		os.Exit(1)
	} else {
		slog.Info("Migrations applied successfully ")
	}

	slog.Info("Starting server...")
	srv := server.New(cfg, conn)
	if err := srv.Start(); err != nil {
		slog.Error("Failed to start server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
