package models

import "time"


type Goal struct {
	DbModel
	UserId   uint          `json:"user_id"`
	GameId   uint          `json:"game_id"` 
	Game     Game          `gorm:"foreignKey:GameId"` 
	TeamId   uint          `json:"team_id"` 
	Team     Team          `gorm:"foreignKey:TeamId"` 
	Duration time.Duration `json:"duration"`
	// Assist1, Assist2, PP
}