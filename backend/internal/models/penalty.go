package models

type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id" validate:"required"`
	TeamID        uint        `json:"team_id" validate:"required"`
	GameID        uint        `json:"game_id" validate:"required"`
	Period        uint        `json:"period" validate:"required"`
	Duration      uint        `json:"duration" validate:"required"`
	CreatedBy     uint        `json:"created_by" validate:"required"`
	PenaltyType   PenaltyType `json:"penalty_type" validate:"required"`
	PenaltyTypeID uint        `json:"penalty_type_id" validate:"required"`
}
