package models

type Roster struct {
	DbModel
	Players   []*User `json:"players" gorm:"many2many:player_rosters"`
	Captain   User    `json:"captain" validate: "required"`
	CaptainID uint    `json:"captain_id" validate: "required"`
}
