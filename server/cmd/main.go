package main

import (
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server"
	"github.com/jak103/powerplay/internal/utils/log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to load config")
		return
	}

	err = log.Init(config.Vars.ColorLog)
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize logger")
		return
	}

	log.Info("--Power Play v0.0.0--") // TODO Get build info automatically

	// Connect to DB
	err = db.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to connect to DB")
		return
	}

	err = db.Migrate()
	if err != nil {
		log.WithErr(err).Alert("Migrations failed")
		return
	}

	// init server
	err = server.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize server")
		return
	}

	// run
	server.Run()
}
