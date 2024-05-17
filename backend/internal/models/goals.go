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
	Assist1Id	uint   `json:"user_id"`
	Assist2Id	uint   `json:"user_id"`
	Powerplay	bool	`json:"powerplay"`
	Penatly		bool	`json:"penalty"`

	//powerplay - was someone in the box; bool
	//penalty - was scored on penalty shot; bool
}