package main

import (
	"github.com/leagueify/league/internal/config"
	"github.com/leagueify/league/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewEchoServer(cfg).Start()
}
