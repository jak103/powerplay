package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"time"
)

type response struct {
    TeamStats []teamData `json:"teamStats"`
}

type teamData struct {
    Name                    string                  `json:"name"`
	League                  string                  `json:"league"`
	EarlyGames              int                     `json:"earlyGame"`
	LateGames               int                     `json:"lateGame"`
	DaysOfTheWeek           map[time.Weekday]int    `json:"daysOfTheWeek"`
	DaysBetweenGames        []int                   `json:"daysBetweenGames"`
	AverageDaysBetweenGames float32                 `json:"averageDaysBetweenGames"`
	Games                   []game                  `json:"games"`
	Balanced                bool                    `json:"balanced"`
}

type game struct {
    Start       string    `json:"start"`
	StartDate   string    `json:"startDate"`
	StartTime   string    `json:"startTime"`
	EndDate     string    `json:"endDate"`
	EndTime     string    `json:"endTime"`
	Location    string    `json:"location"`     // George S. Eccles Ice Center --- Surface 1
	LocationUrl string    `json:"locationUrl"` // https://www.google.com/maps?cid=12548177465055817450
	EventType   string    `json:"eventType"`   // Must be "Game"
	League      string    `json:"league"`            // Not in CSV
	Team1Id     string    `json:"team1Id"`
	Team2Id     string    `json:"team2Id"`
	Team1Name   string    `json:"team1Name"`
	Team2Name   string    `json:"team2Name"`
	IsEarly     bool      `json:"isEarly"`
	Optimized   bool      `json:"optimized"`
}

func init() {
    apis.RegisterHandler(fiber.MethodGet, "/schedule/analysis/:scheduleId<int/>", auth.ManagerOnly, handleAnalysis)
}

func handleAnalysis(c *fiber.Ctx) error {
	games, seasonConfig := parser.ReadGames("summer_2024")

    scheduleId := c.Params("scheduleId")
    //TODO: get the schedule from the database with this id
    log.Info(scheduleId)

    // TODO: get team stats and serialize all of the teams stats
    // Research if we are storing these in a database, if not, we can store them
    // The map contains the team name, an example json object might look like this:
    // {
    //     teamName1: {
    //         team1Data
    //     }
    //     teamName2: {
    //         team2Data
    //     }
    // }
    _, ts := analysis.RunTimeAnalysis(games)

	printTeamSchedules(games, seasonConfig)

    data := response {
        TeamStats: serialize(ts),
    }

	return responder.OkWithData(c, data)
}

func serialize(ts map[string]models.TeamStats) []teamData {
    var stats []teamData
    for _, v := range ts {
        var games []game
        for _, g := range v.Games {
            game := game {
                Start: g.Start.String(),
                StartDate: g.StartDate,
                StartTime: g.StartTime,
                EndDate: g.EndDate,
                EndTime: g.EndTime,
                Location: g.Location,
                LocationUrl: g.LocationUrl,
                Team1Id: g.Team1Id,
                Team2Id: g.Team2Id,
                Team1Name: g.Team1Name,
                Team2Name: g.Team2Name,
                IsEarly: g.IsEarly,
                Optimized: g.Optimized,
            }
            games = append(games, game)
        }
        td := teamData {
            Name: v.Name,
            League: v.League,
            EarlyGames: v.EarlyGames,
            LateGames: v.LateGames,
            DaysOfTheWeek: v.DaysOfTheWeek,
            DaysBetweenGames: v.DaysBetweenGames,
            AverageDaysBetweenGames: v.AverageDaysBetweenGames,
            Games: games,
            Balanced: v.Balanced,
        }
        stats = append(stats, td)
    }
    return stats
}

func printTeamSchedules(games []models.Game, seasonConfig models.SeasonConfig) {
	for _, league := range seasonConfig.Leagues {
		for _, team := range league.Teams {
			log.Info("-----------\n%v\n", team.Name)
			for _, game := range games {
				if team.Name == game.Team1Name || team.Name == game.Team2Name {
					log.Info("%s\n", game)
				}
			}
		}
	}
}
