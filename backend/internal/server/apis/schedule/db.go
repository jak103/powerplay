package schedule

import (
	"gorm.io/gorm"
)

type IceTime struct {
	ID        int    `json:"id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type TeamSchedule struct {
	ID       int       `json:"id"`
	TeamID   int       `json:"team_id"`
	TeamName string    `json:"team_name"`
	IceTimes []IceTime `json:"ice_times"`
}

type LeagueSchedule struct {
	ID            int            `json:"id"`
	LeagueID      int            `json:"league_id"`
	LeagueName    string         `json:"league_name"`
	TeamSchedules []TeamSchedule `json:"team_schedules"`
}

type SeasonSchedule struct {
	ID              int              `json:"id"`
	SeasonID        int              `json:"season_id"`
	SeasonName      string           `json:"season_name"`
	LeagueSchedules []LeagueSchedule `json:"league_schedules"`
}

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
func InsertSeasonSchedule(db *gorm.DB, seasonSchedule SeasonSchedule) (int, error) {
	result := db.Create(&seasonSchedule)
	return seasonSchedule.ID, result.Error
}
