package models

import (
	"fmt"
	"time"
)

// TODO need to move these into the models package in the internal directory
// TODO need to update the models I created based off of these

type Season struct {
	LeagueRounds map[string][]Round
}

type Round struct {
	Games []Game
}

// https://help.sportsengine.com/en/articles/6380725-schedule-upload-tutorial
// https://intercom.help/SportsEngine/en/articles/6310600-schedule-upload-fields-reference-guide
type Game struct {
	Start       time.Time `csv:"-"`
	StartDate   string    `csv:"Start_Date"`
	StartTime   string    `csv:"Start_Time"`
	End         time.Time `csv:"-"`
	EndDate     string    `csv:"End_Date"`
	EndTime     string    `csv:"End_Time"`
	Location    string    `csv:"Location"`     // George S. Eccles Ice Center --- Surface 1
	LocationUrl string    `csv:"Location_URL"` // https://www.google.com/maps?cid=12548177465055817450
	EventType   string    `csv:"Event_Type"`   // Must be "Game"
	League      string    `csv:"-"`            // Not in CSV
	Team1Id     string    `csv:"Team1_ID"`
	Team2Id     string    `csv:"Team2_ID"`
	Team1Name   string    `csv:"Team1_Name"`
	Team2Name   string    `csv:"Team2_Name"`
	IsEarly     bool      `csv:"-"`
	Optimized   bool      `csv:"-"`
}

func (g Game) String() string {
	return fmt.Sprintf("%s %s v %s", g.Start.Format("01-02-2006 03:04"), g.Team1Name, g.Team2Name)
	// DateOnly   = "2006-01-02"
	// TimeOnly   = "15:04:05"
}

// Config
type SeasonConfig struct {
	Leagues  []League `yaml:"leagues"`
	IceTimes []string `yaml:"ice_time"`
}

type League struct {
	Name  string `yaml:"league"`
	Teams []Team `yaml:"teams"`
}

type Team struct {
	Name string `yaml:"name"`
	Id   string `yaml:"id"`
}

type LeagueIceTimes struct {
	Name  string   `yaml:"league"`
	Times []string `yaml:"times"`
	Used  []bool   `yaml:"-"`
}

type TeamStats struct {
	Name                    string
	League                  string
	EarlyGames              int
	LateGames               int
	DaysOfTheWeek           map[time.Weekday]int
	DaysBetweenGames        []int
	AverageDaysBetweenGames float32
	Games                   []Game
	Balanced                bool
}

type SeasonStats struct {
	EarlyGames int
	LateGames  int
	TotalGames int
}

func (ss SeasonStats) EarlyPercentage() float32 {
	return (float32(ss.EarlyGames) / float32(ss.TotalGames))
}
