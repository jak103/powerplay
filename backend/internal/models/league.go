package models

type LeagueRecord struct {
	DbModel
	LeagueID string `json:"league_id"`
	SeasonID uint   `json:"season_id"`
	Name     string `json:"name"`
	Teams    []Team `json:"teams"`
}
