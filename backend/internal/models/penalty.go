package models

type PenaltyType struct {
	DbModel
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
	Severity string `json:"severity"`
}

type Penalty struct {
	DbModel
	PlayerID      uint        `json:"player_id" gorm:"not_null"`
	TeamID        uint        `json:"team_id" gorm:"not_null"`
	GameID        uint        `json:"game_id" gorm:"not_null"`
	Period        uint        `json:"period" gorm:"not_null"`
	Duration      uint        `json:"duration"`
	CreatedBy     uint        `json:"created_by" gorm:"not_null"`
	PenaltyType   PenaltyType `json:"penalty_type"`
	PenaltyTypeID uint        `json:"penalty_type_id" gorm:"not_null"`
}
