package models

type Goal struct {
	DbModel
	UserId uint `json:"user_id"`
	GameId uint `json:"game_id"`
	//Game     			Game          	`gorm:"game"` TODO: When seeding is finished we can officially test this
	TeamId uint `json:"team_id"`
	//Team     			Team			`gorm:"team"` TODO: When seeding is finished we can officially test this
	Duration           uint `json:"duration"` // Do we potentially want to change this to time.Duration
	Period             uint `json:"period"`
	Assist1Id          uint `json:"assist1_id"`
	Assist2Id          uint `json:"assist2_id"`
	PlayerDifferential int  `json:"playerdifferential"`
	IsPenaltyShot      bool `json:"ispenaltyshot"`

	//powerplay - was someone in the box; bool
	//penalty - was scored on penalty shot; bool
}
