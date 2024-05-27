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
	HomeTeamID          uint   `json:"home_team_id"`
	HomeTeamRoster      Roster `json:"home_team_roster"`
	HomeTeamRosterID    uint   `json:"home_team_roster_id"`
	HomeTeamLockerRoom  string `json:"home_team_locker_room"`
	HomeTeamShotsOnGoal int    `json:"home_team_shots_on_goal"`
	HomeTeamScore       int    `json:"home_team_score"`

	AwayTeam            Team   `json:"away_team"`
	AwayTeamID          uint   `json:"away_team_id"`
	AwayTeamRoster      Roster `json:"away_team_roster"`
	AwayTeamRosterID    uint   `json:"away_team_roster_id"`
	AwayTeamLockerRoom  string `json:"away_team_locker_room"`
	AwayTeamShotsOnGoal int    `json:"away_team_shots_on_goal"`
	AwayTeamScore       int    `json:"away_team_score"`

	ScoreKeeper        *User `json:"score_keeper"`
	ScoreKeeperID      uint  `json:"score_keeper_id"`
	PrimaryReferee     *User `json:"primary_referee"`
	PrimaryRefereeID   *uint `json:"primary_referee_id"`
	SecondaryReferee   *User `json:"secondary_referee"`
	SecondaryRefereeID *uint `json:"secondary_referee_id"`
}
