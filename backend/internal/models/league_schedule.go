package models

type LeagueSchedule struct {
	DbModel
	LeagueID      int            `json:"league_id"`
	LeagueName    string         `json:"league_name"`
	TeamSchedules []TeamSchedule `json:"team_schedules"`
}
