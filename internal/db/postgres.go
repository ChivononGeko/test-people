package db

import (
	"context"
	"test-people/internal/config"
	"time"

	"github.com/jackc/pgx/v4"
)

func NewPostgresConnection(cfg config.Config) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
