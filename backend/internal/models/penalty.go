package models

import "time"

type PenaltyType struct {
	DbModel
	Name     string        `json:"name"`
	PlayerID uint          `json:"player_id"`
	Player   User          `json:"player"`
	Duration time.Duration `json:"duration"`
	Severity string        `json:"severity"`
}

type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id"`
	Player        User        `json:"player"`
	TeamID        uint        `json:"team_id"`
	Team          Team        `json:"team"`
	GameID        uint        `json:"game_id"`
	Game          Game        `json:"game"`
	Period        uint        `json:"period"` // TODO: how to represent Shootouts?
	Duration      uint        `json:"duration"`
	CreatedBy     uint        `json:"created_by"` // TODO: Should this be in the model.go DbModel?
	PenaltyTypeID uint        `json:"penalty_type_id"`
	PenaltyType   PenaltyType `json:"penalty_type"`
}
