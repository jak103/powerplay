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
	SeasonID uint      `json:"season_id" validate:"required"`
	Start    time.Time `json:"start" validate:"required"`
	Venue    *Venue     `json:"venue"`
	VenueID  *uint      `json:"venue_id" validate:"required"`
	Status   Status    `json:"status" validate:"required"`

	HomeTeam            *Team   `json:"home_team"`
	HomeTeamID          uint   `json:"home_team_id" validate:"required"`
	HomeTeamRoster      *Roster `json:"home_team_roster"`
	HomeTeamRosterID    *uint   `json:"home_team_roster_id" validate:"required"`
	HomeTeamLockerRoom  string `json:"home_team_locker_room" validate:"required"`
	HomeTeamShotsOnGoal int    `json:"home_team_shots_on_goal" validate:"required"`
	HomeTeamScore       int    `json:"home_team_score" validate:"required"`

	AwayTeam            *Team   `json:"away_team"`
	AwayTeamID          uint   `json:"away_team_id" validate:"required"`
	AwayTeamRoster      *Roster `json:"away_team_roster"`
	AwayTeamRosterID    *uint   `json:"away_team_roster_id" validate:"required"`
	AwayTeamLockerRoom  string `json:"away_team_locker_room" validate:"required"`
	AwayTeamShotsOnGoal int    `json:"away_team_shots_on_goal" validate:"required"`
	AwayTeamScore       int    `json:"away_team_score" validate:"required"`

	ScoreKeeper        *User `json:"score_keeper"`
	ScoreKeeperID      uint  `json:"score_keeper_id" validate:"required"`
	PrimaryReferee     *User `json:"primary_referee"`
	PrimaryRefereeID   *uint `json:"primary_referee_id" validate:"required"`
	SecondaryReferee   *User `json:"secondary_referee"`
	SecondaryRefereeID *uint `json:"secondary_referee_id" validate:"required"`
}
