package schedule

import "gorm.io/gorm"

type IceTime struct {
	gorm.Model
	ID        int    `yaml:"id"`
	StartTime string `yaml:"start_time"`
}

type TeamSchedule struct {
	gorm.Model
	ID       int       `yaml:"id"`
	TeamID   int       `yaml:"team_id"`
	TeamName string    `yaml:"team_name"`
	IceTimes []IceTime `yaml:"ice_times"`
}

type LeagueSchedule struct {
	gorm.Model
	ID            int            `yaml:"id"`
	LeagueID      int            `yaml:"league_id"`
	LeagueName    string         `yaml:"league_name"`
	TeamSchedules []TeamSchedule `yaml:"team_schedules"`
}

type SeasonSchedule struct {
	gorm.Model
	ID              int              `yaml:"id"`
	SeasonID        int              `yaml:"season_id"`
	SeasonName      string           `yaml:"season_name"`
	LeagueSchedules []LeagueSchedule `yaml:"league_schedules"`
}
