package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
	"gorm.io/gorm"
)

var migrations []*gormigrate.Migration

func init() {
	// todo: break into migration files #59 - https://github.com/jak103/powerplay/issues/59
	migrations = append(migrations,
		&gormigrate.Migration{
			ID: "create_penalty_type_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.PenaltyType{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("penalty_types")
			},
		},
		&gormigrate.Migration{
			ID: "penalty_type_remove_player_column",
			Migrate: func(tx *gorm.DB) error {
				if tx.Migrator().HasColumn(&models.PenaltyType{}, "PlayerID") {
					if err := tx.Migrator().DropColumn(&models.PenaltyType{}, "PlayerID"); err != nil {
						log.WithErr(err).Alert("Error dropping column 'PlayerID' on PenaltyType")
						return err
					}
				}
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if !tx.Migrator().HasColumn(&models.PenaltyType{}, "player_id") {
					if err := tx.Migrator().AddColumn(&models.PenaltyType{}, "player_id"); err != nil {
						return err
					}
				}
				return nil
			},
		},
		&gormigrate.Migration{
			ID: "create_goals_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Goal{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("goals")
			},
		},

		&gormigrate.Migration{
			ID: "create_shots_on_goal_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.ShotOnGoal{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("goals")
			},
		},
		// Add more migrations here
	)
}

func createMigrator(db *gorm.DB, migrations []*gormigrate.Migration) *gormigrate.Gormigrate {
	var migrator *gormigrate.Gormigrate
	if migrations != nil {
		migrator = gormigrate.New(db, &gormigrate.Options{
			TableName:                 "migrations",
			IDColumnName:              "id",
			IDColumnSize:              255,
			UseTransaction:            true,
			ValidateUnknownMigrations: false,
		}, migrations)
	} else {
		empty_migration := []*gormigrate.Migration{}
		migrator = gormigrate.New(db, &gormigrate.Options{
			TableName:                 "migrations",
			IDColumnName:              "id",
			IDColumnSize:              255,
			UseTransaction:            true,
			ValidateUnknownMigrations: false,
		}, empty_migration)
	}
	return migrator
}

func Run(db *gorm.DB) error {
	log.Info("Running migrations")

	initialMigrator := createMigrator(db, nil)
	// auatomigrate all objects
	initialMigrator.InitSchema(func(tx *gorm.DB) error {
		log.Info("Initializing powerplay schema")
		err := tx.AutoMigrate(
			&models.User{},
			&models.Season{},
			&models.League{},
			&models.Team{},
			&models.Roster{},
			&models.Venue{},
			&models.Game{},
			&models.Registration{},
			&models.KeyRecord{},
			&models.Goal{},
			&models.PenaltyType{},
			&models.Penalty{},
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
	err := initialMigrator.Migrate()
	if err != nil {
		log.WithErr(err).Alert("Failed to run migrations")
		return err
	}

	additionalMigrator := createMigrator(db, migrations)
	err = additionalMigrator.Migrate()
	if err != nil {
		log.WithErr(err).Alert("Failed to run migrations")
		return err
	}

	log.Info("Migrations complete")
	return nil
}
