package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/league/handler"
	"github.com/leagueify/league/internal/config"
	"github.com/leagueify/league/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	leagueRouter := handler.LeagueRouter()

	router.Handle("/league/", http.StripPrefix("/league", leagueRouter))

	// define server config
	leagueServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := leagueServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
