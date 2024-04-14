package db

import "github.com/jak103/powerplay/internal/utils/log"

func (s session) ResetDatabase() error {
	result := db.Raw("DROP DATABASE powerplay WITH (FORCE)")
	if result.Error != nil {
		log.WithErr(result.Error).Error("Failed to drop database")
		return result.Error
	}

	log.Debug("%v rows affected", result.RowsAffected)

	err := Migrate()
	if err != nil {
		log.WithErr(err).Error("Failed to migrate database")
		return err
	}

	return nil
}
