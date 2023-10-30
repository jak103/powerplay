package main

import (
	"github.com/jak103/leaguemanager/internal/config"
	"github.com/jak103/leaguemanager/internal/log"
	"github.com/jak103/leaguemanager/internal/server"
)

func main() {
	log.Info("--League Manager v0.0.0--") // TODO Get build info automatically

	// load env config
	cfg, err := config.Load()
	if err != nil {
		// log.WithErr(err).Error()
	}

	// init logger
	err = log.Init(cfg)
	if err != nil {
		// log.WithErr(err).Error()
	}

	// init server
	err = server.Init(cfg)
	if err != nil {
		// log.WithErr(err).Error()
	}

	// run
	server.Run()
}
