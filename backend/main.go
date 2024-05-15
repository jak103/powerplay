package main

import (
	"flag"
	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/server"
	"github.com/jak103/powerplay/internal/utils/log"
)

func runMigrations() {
	err := db.Migrate()
	if err != nil {
		log.WithErr(err).Alert("Migrations failed")
		return
	}
	log.Info("Migrations completed successfully")
}

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations and exit")
	flag.Parse()

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

	// Connect to DB
	err = db.Init()
	if err != nil {
		log.WithErr(err).Alert("Failed to connect to DB")
		return
	}

	if *migrateFlag {
		runMigrations()
		return
	}

	runMigrations()

	// run
	server.Run()
}
