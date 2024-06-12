package main

import (
	"flag"
	"github.com/jak103/powerplay/internal/db/seeders/fake_data"
	"github.com/jak103/powerplay/internal/models"

	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/db"
	ppseeders "github.com/jak103/powerplay/internal/db/seeders"
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

func runSeeds() {
	seeders := []ppseeders.Seeder{
		ppseeders.PenaltyTypeSeeder{},
		// Add more seeders here
	}

	// Seed the database
	if err := db.RunSeeders(seeders); err != nil {
		log.WithErr(err).Alert("Failed to seed database: %v", err)
	}
	log.Info("Successfully seeded database")
}

func runFakeDataSeeds() {
	// Initialize seeders
	seasonSeeder := fake_data.SeasonSeeder{}
	leagueSeeder := fake_data.LeagueSeeder{}
	teamSeeder := fake_data.TeamSeeder{}
	venueSeeder := fake_data.VenueSeeder{}

	// Seed Season
	season, err := seasonSeeder.Seed(db.GetDB())
	if err != nil {
		log.WithErr(err).Alert("Failed to seed Season: %v", err)
		return
	}
	log.Info("Successfully seeded Season")

	// Seed LeagueRecord with SeasonID
	leagues, err := leagueSeeder.Seed(db.GetDB(), season.(models.Season).ID)
	if err != nil {
		log.WithErr(err).Alert("Failed to seed LeagueRecord: %v", err)
		return
	}
	log.Info("Successfully seeded LeagueRecord")

	// Seed Teams for each LeagueRecord
	for _, league := range leagues.([]models.LeagueRecord) {
		_, err = teamSeeder.Seed(db.GetDB(), league.ID)
		if err != nil {
			log.WithErr(err).Alert("Failed to seed Teams for LeagueRecord %d: %v", league.ID, err)
			return
		}
		log.Info("Successfully seeded Teams for LeagueRecord %d", league.ID)
	}

	// Seed Venues
	_, err = venueSeeder.Seed(db.GetDB())
	if err != nil {
		log.WithErr(err).Alert("Failed to seed Venues: %v", err)
		return
	}
	log.Info("Successfully seeded Venues")
}

func main() {
	migrateFlag := flag.Bool("migrate", false, "Run database migrations and exit")
	seedTestData := flag.Bool("seed-test", false, "Seed test data and exit")
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

	if *seedTestData {
		runFakeDataSeeds()
		return
	}

	runMigrations()
	runSeeds()

	// run
	server.Run()
}
