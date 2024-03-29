package main

import (
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/server"
	"github.com/jak103/powerplay/internal/utils/log"
)

func main() {
	err := log.Init("DEBUG", false)
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize logger")
		return
	}

	err = config.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to load config")
		return
	}

	err = log.Init(config.Vars.LogLevel, config.Vars.LogColor)
	if err != nil {
		log.WithErr(err).Alert("Failed to initialize logger")
		return
	}

	log.Info("--Power Play v0.0.0--") // TODO Get build info automatically

	// TODO Don't need the DB for the push tests, renable later
	// Connect to DB
	// err = db.Init()
	// if err != nil {
	// 	log.WithErr(err).Alert("Failed to connect to DB")
	// 	return
	// }

	// err = db.Migrate()
	// if err != nil {
	// 	log.WithErr(err).Alert("Migrations failed")
	// 	return
	// }

	// run
	server.Run()
}
