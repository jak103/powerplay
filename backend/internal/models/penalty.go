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
