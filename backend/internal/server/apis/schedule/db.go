package schedule

import (
	_ "github.com/jak103/powerplay/internal/models"
	"gorm.io/gorm"
)

// TODO - fix the models package import

// GetIceTimes retrieves all ice times from the database
func GetIceTimes(db *gorm.DB) ([]IceTime, error) {
	var iceTimes []IceTime
	result := db.Find(&iceTimes)
	return iceTimes, result.Error
}

// CreateSeasonScheduleTable is not needed with GORM, as AutoMigrate handles it
func CreateSeasonScheduleTable(db *gorm.DB) error {
	return db.AutoMigrate(&SeasonSchedule{})
}

// InsertSeasonSchedule inserts a new season schedule into the database
func InsertSeasonSchedule(db *gorm.DB, seasonSchedule SeasonSchedule) (uint, error) {
	result := db.Create(&seasonSchedule)
	return seasonSchedule.ID, result.Error
}
