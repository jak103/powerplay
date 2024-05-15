package models

type SeasonSchedule struct {
	DbModel
	SeasonID        int              `json:"season_id"`
	SeasonName      string           `json:"season_name"`
	LeagueSchedules []LeagueSchedule `json:"league_schedules"`
}
