package models

type Goal struct {
	DbModel
	UserId   			uint          	`json:"user_id" validate:"required"`
	GameId   			uint          	`json:"game_id" validate:"required"` 
	//Game     			Game          	`gorm:"game"` TODO: When seeding is finished we can officially test this
	TeamId   			uint          	`json:"team_id" validate:"required"` 
	//Team     			Team			`gorm:"team"` TODO: When seeding is finished we can officially test this
	Duration 			uint			`json:"duration" validate:"required"` // Do we potentially want to change this to time.Duration
	Period	    		uint			`json:"period" validate:"required"`
	Assist1Id			uint   	   		`json:"assist1_id" validate:"required"`
	Assist2Id			uint   	   		`json:"assist2_id" validate:"required"`
	PlayerDifferential	int		   		`json:"playerdifferential" validate:"required"`
	IsPenaltyShot		bool	   		`json:"ispenaltyshot" validate:"required"`

	//powerplay - was someone in the box; bool
	//penalty - was scored on penalty shot; bool
}
