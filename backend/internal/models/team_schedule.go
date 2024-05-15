package models

type TeamSchedule struct {
	DbModel
	TeamID   int       `json:"team_id"`
	TeamName string    `json:"team_name"`
	IceTimes []IceTime `json:"ice_times"`
}
