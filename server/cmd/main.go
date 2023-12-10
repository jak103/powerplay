package main

import (
	"github.com/jak103/leaguemanager/internal/config"
	"github.com/jak103/leaguemanager/internal/server"
	"github.com/jak103/leaguemanager/internal/utils/log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to load config")
		return
	}

	err = log.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize logger")
		return
	}

	// TODO Come up with a good name
	log.Info("--League Manager v0.0.0--") // TODO Get build info automatically

	// init server
	err = server.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize server")
	}

	// run
	server.Run()
}
