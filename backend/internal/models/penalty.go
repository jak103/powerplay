package models

type PenaltyType struct {
	DbModel
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
	Severity string `json:"severity"`
}

type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id"`
	TeamID        uint        `json:"team_id"`
	GameID        uint        `json:"game_id"`
	Period        uint        `json:"period"`
	Duration      uint        `json:"duration"`
	CreatedBy     uint        `json:"created_by"`
	PenaltyType   PenaltyType `json:"penalty_type"`
	PenaltyTypeID uint        `json:"penalty_type_id"`
}
