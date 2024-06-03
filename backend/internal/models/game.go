package models

import "time"

type Game struct {
	DbModel
	AwayTeam            Team      `json:"away_team"`
	AwayTeamID          uint      `json:"away_team_id"`
	AwayTeamLockerRoom  string    `json:"away_team_locker_room"`
	AwayTeamShotsOnGoal uint      `json:"away_team_shots_on_goal"`
	End                 time.Time `json:"end"`
	HomeTeam            Team      `json:"home_team"`
	HomeTeamID          uint      `json:"home_team_id"`
	HomeTeamLockerRoom  string    `json:"home_team_locker_room"`
	HomeTeamShotsOnGoal uint      `json:"home_team_shots_on_goal"`
	LeagueID            uint      `json:"league_id"`
	League              League    `json:"league"`
	ManagerOnCall       *User     `json:"manager_on_call"`
	ManagerOnCallID     *uint
	PrimaryReferee      *User `json:"primary_referee"`
	PrimaryRefereeID    uint
	ScoreKeeper         *User `json:"score_keeper"`
	ScoreKeeperID       *uint
	SecondaryReferee    *User `json:"secondary_referee"`
	SecondaryRefereeID  *uint
	Season              Season    `json:"season_id"`
	SeasonID            uint      `json:"season"`
	Start               time.Time `json:"start"`
	Venue               Venue     `json:"venue"`
	VenueID             uint
	HomeTeamRosterID    uint   `json:"home_team_roster_id"`
	AwayTeamGoals       uint   `json:"away_team_goals"`
	AwayTeamRosterID    uint   `json:"away_team_roster_id"`
	AwayTeamRoster      Roster `json:"away_team_roster"`
}
