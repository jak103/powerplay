package models

import "time"

type PenaltyType struct {
	DbModel
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
	Severity string        `json:"severity"`
}

type Penalty struct {
	DbModel
	PlayerID      uint `json:"player_id"`
	TeamID        uint `json:"team_id"`
	GameID        uint `json:"game_id"`
	Period        uint `json:"period"` // TODO: how to represent Shootouts?
	Duration      uint `json:"duration"`
	CreatedBy     uint `json:"created_by"` // TODO: Should this be in the model.go DbModel?
	PenaltyTypeID uint `json:"penalty_type_id"`
}
