package models

import "time"

type Status string

const (
	SCHEDULED   Status = "Scheduled"
	IN_PROGRESS Status = "In Progress"
	FINAL       Status = "Final"
)

type Game struct {
	DbModel
	SeasonID uint      `json:"season_id"`
	Start    time.Time `json:"start"`
	Venue    Venue     `json:"venue"`
	VenueID  uint      `json:"venue_id"`
	Status   Status    `json:"status"`

	HomeTeam            Team   `json:"home_team"`
	HomeTeamRoster      Roster `json:"home_team_roster"`
	HomeTeamLockerRoom  string `json:"home_team_locker_room"`
	HomeTeamShotsOnGoal int    `json:"home_team_shots_on_goal"`
	HomeTeamScore       int    `json:"home_team_score"`

	AwayTeam            Team   `json:"away_team"`
	AwayTeamRoster      Roster `json:"away_team_roster"`
	AwayTeamLockerRoom  string `json:"away_team_locker_room"`
	AwayTeamShotsOnGoal int    `json:"away_team_shots_on_goal"`
	AwayTeamScore       int    `json:"away_team_score"`

	ScoreKeeper        *User `json:"score_keeper"`
	ScoreKeeperID      uint  `json:"score_keeper_id"`
	PrimaryReferee     *User `json:"primary_referee"`
	PrimaryRefereeID   *User `json:"primary_referee_id"`
	SecondaryReferee   *User `json:"secondary_referee"`
	SecondaryRefereeID *User `json:"secondary_referee_id"`
}
