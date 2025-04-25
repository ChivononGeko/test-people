package config

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	Postgres string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	pg := os.Getenv("POSTGRES_URL")
	if pg == "" {
		log.Fatal("POSTGRES_URL is required")
	}

	return Config{
		Port:     port,
		Postgres: pg,
	}
}
