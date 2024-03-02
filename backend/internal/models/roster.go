package models

type Roster struct {
	DbModel
	SeasonID uint
	Team     Team `json:"team"`
	TeamID   uint
	Season   Season  `json:"season"`
	Players  []*User `json:"players" gorm:"many2many:player_rosters"`
	Staff    []*User `json:"staff" gorm:"many2many:staff_rosters"`
}
