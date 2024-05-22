package models


type ShotOnGoal struct{
	DbModel
	GameId uint `json:"game_id" gorm:"not_null"`
	TeamId uint `json:"team_id" gorm:"not_null"`
	ShotTime uint `json:"shot_time" gorm:"not_null"`
	Scorekeeper uint `json:"scorekeeper gorm:"not_null"`
}

// Should overide GOs incorrect pluralization
func (ShotOnGoal) TableName() string {
    return "shots_on_goal"
}