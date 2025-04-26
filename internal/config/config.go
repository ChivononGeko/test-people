package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port     string
	Postgres string
}

func Load() (*Config, error) {
	port := getEnv("PORT", "8080")

	user, err := mustGetEnv("DB_USER")
	if err != nil {
		return nil, fmt.Errorf("DB_USER is not defined")
	}

	pass, err := mustGetEnv("DB_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("DB_PASSWORD is not defined")
	}
	host, err := mustGetEnv("DB_HOST")
	if err != nil {
		return nil, fmt.Errorf("DB_HOST is not defined")
	}
	dbPort, err := mustGetEnv("DB_PORT")
	if err != nil {
		return nil, fmt.Errorf("DB_PORT is not defined")
	}
	name, err := mustGetEnv("DB_NAME")
	if err != nil {
		return nil, fmt.Errorf("DB_NAME is not defined")
	}
	sslmode := getEnv("DB_SSLMODE", "disable")

	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, pass, host, dbPort, name, sslmode)

	return &Config{
		Port:     port,
		Postgres: pgURL,
	}, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func mustGetEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("%s is required", key)
	}
	return val, nil
}
