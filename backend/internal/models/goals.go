package models

import "time"


type Goal struct {
	DbModel
	UserId   	uint          	`json:"user_id"`
	GameId   	uint          	`json:"game_id"` 
	//Game     	Game          	`gorm:"game"` 
	TeamId   	uint          	`json:"team_id"` 
	//Team     	Team			`gorm:"team"` 
	Duration 	time.Duration	`json:"duration"`
	Period	    uint			`json:"period"`
	Assist1Id	uint   	   		`json:"assist1_id"`
	Assist2Id	uint   	   		`json:"assist2_id"`
	Powerplay	int		   		`json:"powerplay"`
	Penatly		bool	   		`json:"penalty"`

	//powerplay - was someone in the box; bool
	//penalty - was scored on penalty shot; bool
}