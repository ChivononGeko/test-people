package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dsn string, migrationsPath string) error {
	fmt.Println("Using DSN:", dsn)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		dsn,
	)
	if err != nil {
		return fmt.Errorf("create migrate: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("run migrate up: %w", err)
	}

	return nil
}
