package schedule

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/csv"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/optimize"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
	"gorm.io/gorm"
	"time"
)

func HandleGenerate(c *fiber.Ctx) error {
	log.Info("Scheduler v0.1\n")

	log.Info("Reading body\n")
	seasonFileName, numberOfGamesPerTeam, err := ReadBody(c)
	if err != nil {
		log.Error("Error reading body: %v\n", err)
		return err
	}

	log.Info("Reading config file for season\n")
	seasonConfig, err := parser.SeasonConfig(seasonFileName)
	if err != nil {
		log.Error("Error reading file: %v\n", err)
		return responder.BadRequest(c, "Error reading file")
	}

	log.Info("Generating games\n")
	season := GenerateGames(seasonConfig.Leagues, numberOfGamesPerTeam)

	log.Info("Assigning ice times\n")
	games := AssignTimes(seasonConfig.IceTimes, season)

	log.Info("Optimizing schedule\n")
	OptimizeSchedule(games)

	log.Info("Writing csv\n")
	err = csv.GenerateCsv(games, "schedule.csv")
	if err != nil {
		log.Error("Error writing csv: %v\n", err)
		return responder.InternalServerError(c)
	}

	// TODO should be moved to pkg
	log.Info("Saving to database\n")
	db := c.Locals("db").(*gorm.DB)
	for _, game := range games {
		err = db.Create(&game).Error
		if err != nil {
			log.Error("Error saving game to database: %v\n", err)
			return responder.InternalServerError(c)
		}
	}

	return responder.Ok(c, "Schedule generated at schedule.csv and saved to database")
}

func ReadBody(c *fiber.Ctx) (string, int, error) {
	type BodyDto struct {
		SeasonFileName       string `json:"seasonFileName"`
		NumberOfGamesPerTeam int    `json:"numberOfGamesPerTeam"`
	}
	body := c.Body()
	var bodyDto BodyDto
	err := json.Unmarshal(body, &bodyDto)
	if err != nil {
		return "", 0, responder.BadRequest(c, "Error reading body")
	}

	return bodyDto.SeasonFileName, bodyDto.NumberOfGamesPerTeam, nil
}

func OptimizeSchedule(games []models.Game) {
	log.Info("Pre-optimization analysis")
	seasonStats, teamStats := analysis.RunTimeAnalysis(games)

	// Need to make sure games are balanced in
	// - Early / late
	// - Days between games
	balanceCount := GetBalanceCount(&teamStats)
	lastBalanceCount := -1

	for count := 0; balanceCount != lastBalanceCount && count < 25; count++ {
		optimize.Schedule(games, seasonStats, teamStats)

		log.Info("Post-optimization analysis")
		seasonStats, teamStats = analysis.RunTimeAnalysis(games)

		lastBalanceCount = balanceCount
		balanceCount := GetBalanceCount(&teamStats)

		log.Info("Balanced count: %v\n", balanceCount)
	}
}

func GenerateGames(leagues []models.League, numberOfGamesPerTeam int) models.Season {
	season := models.Season{LeagueRounds: make(map[string][]models.Round)}

	for _, league := range leagues {
		numTeams := len(league.Teams)

		// Figure out how many rounds we need to run to get each team the number of games per season
		numberOfGamesPerTeam += ((numTeams * numberOfGamesPerTeam) - (numTeams/2)*(2*numberOfGamesPerTeam)) / 2

		log.Info("League %v games per round: %v\n", league.Name, numberOfGamesPerTeam)

		if numTeams%2 == 1 {
			league.Teams = append(league.Teams, models.Team{Name: "Bye", Id: "-1"})
			numTeams = len(league.Teams)
		}

		numberOfRounds := numberOfGamesPerTeam

		rounds := make([]models.Round, numberOfRounds)

		for round := 0; round < numberOfRounds; round++ {
			rounds[round].Games = make([]models.Game, numTeams/2)
			for i := 0; i < numTeams/2; i++ {
				team1 := league.Teams[i].Id
				team1Name := league.Teams[i].Name
				team2 := league.Teams[numTeams-1-i].Id
				team2Name := league.Teams[numTeams-1-i].Name

				rounds[round].Games[i] = NewGame(league.Name, team1, team1Name, team2, team2Name)
			}

			RotateTeams(&league)
		}
		season.LeagueRounds[league.Name] = rounds
	}

	return season
}

func AssignTimes(times []string, season models.Season) []models.Game {

	games := NewGames(&season)

	log.Info("Have times for %v games\n", len(times))
	log.Info("Have %v games\n", len(games))
	for i := range games {
		startTime, err := time.Parse("1/2/06 15:04", times[i])
		if err != nil {
			log.Error("Failed to parse start time: %v\n", err)
		}
		endTime := startTime.Add(75 * time.Minute)

		games[i].Start = startTime
		games[i].StartDate = startTime.Format("01/02/2006")
		games[i].StartTime = startTime.Format("15:04")

		games[i].EndDate = endTime.Format("01/02/2006")
		games[i].EndTime = endTime.Format("15:04")

		switch games[i].Start.Hour() {
		case 20:
			games[i].IsEarly = true
		case 21:
			if games[i].Start.Minute() <= 15 {
				games[i].IsEarly = true
			} else {
				games[i].IsEarly = false
			}
		case 22, 23:
			games[i].IsEarly = false
		}
	}

	return games
}

func GetBalanceCount(teamStats *map[string]models.TeamStats) int {
	balanceCount := 0
	for _, team := range *teamStats {
		if team.Balanced {
			balanceCount++
		}
	}
	return balanceCount
}

func RotateTeams(league *models.League) {
	// Rotate teams except the first one
	lastTeam := league.Teams[len(league.Teams)-1]
	copy(league.Teams[2:], league.Teams[1:len(league.Teams)-1])
	league.Teams[1] = lastTeam
}

func NewGame(league, team1, team1Name, team2, team2Name string) models.Game {
	return models.Game{
		Team1Id:     team1,
		Team1Name:   team1Name,
		Team2Id:     team2,
		Team2Name:   team2Name,
		League:      league,
		Location:    "George S. Eccles Ice Center --- Surface 1",
		LocationUrl: "https://www.google.com/maps?cid=12548177465055817450",
		EventType:   "Game",
	}
}

func NewGames(season *models.Season) []models.Game {
	games := make([]models.Game, 0)
	for i := 0; i < 10; i += 1 { // Rounds // TODO This currently won't work if the leagues don't all have the same number of teams, fix this when needed (Balance by calculating the rate at which games have to be assigned, e.g. the average time between games to complete in the season from the number of first to last dates )
		for _, league := range []string{"A", "C", "B", "D"} { // Alternate leagues so if you play in two leagues you don't play back to back
			for j, game := range season.LeagueRounds[league][i].Games {
				if game.Team1Id != "-1" && game.Team2Id != "-1" { // TODO what does -1 mean?
					games = append(games, season.LeagueRounds[league][i].Games[j])
				}
			}
		}
	}
	return games
}

func IsEarlyGame(hour, minute int) bool {
	switch hour {
	case 20:
		return true
	case 21:
		return minute <= 15
	case 22, 23:
		return false
	}
	return false
}
