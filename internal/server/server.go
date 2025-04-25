package server

import (
	"fmt"
	"net/http"
	"test-people/internal/adapters/enrichment"
	"test-people/internal/adapters/repository"
	"test-people/internal/adapters/transport"
	"test-people/internal/config"
	"test-people/internal/service"

	"github.com/jackc/pgx/v4"
)

type Server struct {
	cfg    config.Config
	db     *pgx.Conn
	router http.Handler
}

func New(cfg config.Config, db *pgx.Conn) *Server {
	personRepo := repository.NewPostgresPersonRepository(db)
	enricher := enrichment.NewExternalDataAdapter()
	service := service.NewPersonService(personRepo, enricher)
	handler := transport.NewPersonHandler(service)

	router := transport.NewRouter(handler)

	return &Server{
		cfg:    cfg,
		db:     db,
		router: router,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.cfg.Port)
	return http.ListenAndServe(addr, s.router)
}
