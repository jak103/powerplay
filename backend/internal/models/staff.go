package models

type Staff struct {
	dbModel
	ManagerOnCall    User `json:"manager_on_call"`
	ScoreKeeper      User `json:"score_keeper"`
	PrimaryReferee   User `json:"primary_referee"`
	SecondaryReferee User `json:"secondary_referee"`
	GameId           int  `json:"-"`
}
