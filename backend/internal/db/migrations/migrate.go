package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/gorm"
)

var migrations []*gormigrate.Migration

func Run(db *gorm.DB) error {
	log.Info("Running migrations")

	migrator := gormigrate.New(db, &gormigrate.Options{
		TableName:                 "migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: false,
	}, migrations)

	// auatomigrate all objects
	migrator.InitSchema(func(tx *gorm.DB) error {
		log.Info("Initializing powerplay schema")
		err := tx.AutoMigrate(
			&models.User{},
			&models.League{},
			&models.Team{},
			&models.Roster{},
			&models.Staff{},
			&models.Game{},
			&models.Season{},
			&models.Registration{},
			&models.Venue{},
			&models.KeyRecord{},
		)
		if err != nil {
			return err
		}

		// if err := tx.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_organizations FOREIGN KEY (organization_id) REFERENCES organizations (id)").Error; err != nil {
		// 	return err
		// }
		// all other constraints, indexes, etc...
		return nil
	})

	err := migrator.Migrate()
	if err != nil {
		log.WithErr(err).Alert("Failed to run migrations")
		return err
	}

	log.Info("Migrations complete")
	return nil
}
