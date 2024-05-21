package schedule

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"

	"github.com/jak103/powerplay/internal/db"
	dbModels "github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/csv"
	scheduleModels "github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/optimize"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/parser"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func handleGenerate(c *fiber.Ctx) error {
	log.Info("Scheduler v0.1\n")

	log.Info("Reading body\n")
	seasonFileName, numberOfGamesPerTeam, err := readBody(c)
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
	season := generateGames(seasonConfig.Leagues, numberOfGamesPerTeam)

	log.Info("Assigning ice times\n")
	games := assignTimes(seasonConfig.IceTimes, season)

	log.Info("Optimizing schedule\n")
	optimizeSchedule(games)

	log.Info("Writing csv\n")
	err = csv.GenerateCsv(games, "schedule.csv")
	if err != nil {
		log.Error("Error writing csv: %v\n", err)
		return responder.InternalServerError(c)
	}

	log.Info("Saving to database\n")
	session := db.GetSession(c)
	err = saveToDb(&session, games)
	if err != nil {
		log.Error("Error saving to database: %v\n", err)
		return responder.InternalServerError(c)
	}

	return responder.Ok(c, "Schedule generated at schedule.csv and saved to database")
}

func saveToDb(session *db.Session, games []models.Game) error {
	dbGames := make([]dbModels.Game, len(games))
	for i, game := range games {
		dbGames[i] = mapScheduleGameToDbGame(game)
	}

	for _, dbGame := range dbGames {
		if err := session.Connection.Save(&dbGame).Error; err != nil {
			return err
		}
	}

	return nil
}

func mapScheduleGameToDbGame(game scheduleModels.Game) dbModels.Game {
	return dbModels.Game{
		Teams: []dbModels.Team{
			{Name: game.Team1Name},
			{Name: game.Team2Name},
		},
		Start: game.Start,
		End:   game.End,
	}
}

func readBody(c *fiber.Ctx) (string, int, error) {
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

func optimizeSchedule(games []scheduleModels.Game) {
	log.Info("Pre-optimization analysis")
	seasonStats, teamStats := analysis.RunTimeAnalysis(games)

	// Need to make sure games are balanced in
	// - Early / late
	// - Days between games
	balanceCount := getBalanceCount(&teamStats)
	lastBalanceCount := -1

	for count := 0; balanceCount != lastBalanceCount && count < 25; count++ {
		optimize.Schedule(games, seasonStats, teamStats)

		log.Info("Post-optimization analysis")
		seasonStats, teamStats = analysis.RunTimeAnalysis(games)

		lastBalanceCount = balanceCount
		balanceCount := getBalanceCount(&teamStats)

		log.Info("Balanced count: %v\n", balanceCount)
	}
}

func generateGames(leagues []scheduleModels.League, numberOfGamesPerTeam int) scheduleModels.Season {
	season := scheduleModels.Season{LeagueRounds: make(map[string][]scheduleModels.Round)}

	for _, league := range leagues {
		numTeams := len(league.Teams)

		// Figure out how many rounds we need to run to get each team the number of games per season
		numberOfGamesPerTeam += ((numTeams * numberOfGamesPerTeam) - (numTeams/2)*(2*numberOfGamesPerTeam)) / 2

		log.Info("League %v games per round: %v\n", league.Name, numberOfGamesPerTeam)

		if numTeams%2 == 1 {
			league.Teams = append(league.Teams, scheduleModels.Team{Name: "Bye", Id: "-1"})
			numTeams = len(league.Teams)
		}

		numberOfRounds := numberOfGamesPerTeam

		rounds := make([]scheduleModels.Round, numberOfRounds)

		for round := 0; round < numberOfRounds; round++ {
			rounds[round].Games = make([]scheduleModels.Game, numTeams/2)
			for i := 0; i < numTeams/2; i++ {
				team1 := league.Teams[i].Id
				team1Name := league.Teams[i].Name
				team2 := league.Teams[numTeams-1-i].Id
				team2Name := league.Teams[numTeams-1-i].Name

				rounds[round].Games[i] = newGame(league.Name, team1, team1Name, team2, team2Name)
			}

			rotateTeams(&league)
		}
		season.LeagueRounds[league.Name] = rounds
	}

	return season
}

func assignTimes(times []string, season scheduleModels.Season) []scheduleModels.Game {

	games := newGames(&season)

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

		games[i].End = endTime
		games[i].EndDate = endTime.Format("01/02/2006")
		games[i].EndTime = endTime.Format("15:04")

		games[i].IsEarly = isEarlyGame(games[i].Start.Hour(), games[i].Start.Minute())
	}

	return games
}

func getBalanceCount(teamStats *map[string]scheduleModels.TeamStats) int {
	balanceCount := 0
	for _, team := range *teamStats {
		if team.Balanced {
			balanceCount++
		}
	}
	return balanceCount
}

func rotateTeams(league *scheduleModels.League) {
	// Rotate teams except the first one
	lastTeam := league.Teams[len(league.Teams)-1]
	copy(league.Teams[2:], league.Teams[1:len(league.Teams)-1])
	league.Teams[1] = lastTeam
}

func newGame(league, team1, team1Name, team2, team2Name string) scheduleModels.Game {
	return scheduleModels.Game{
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

func newGames(season *scheduleModels.Season) []scheduleModels.Game {
	games := make([]scheduleModels.Game, 0)
	for i := 0; i < 10; i += 1 { // Rounds // TODO This currently won't work if the leagues don't all have the same number of teams, fix this when needed (Balance by calculating the rate at which games have to be assigned, e.g. the average time between games to complete in the season from the number of first to last dates )
		for _, league := range []string{"A", "C", "B", "D"} { // Alternate leagues so if you play in two leagues you don't play back to back
			if season.LeagueRounds[league] == nil || len(season.LeagueRounds[league]) <= i {
				continue
			}
			for j, game := range season.LeagueRounds[league][i].Games {
				if game.Team1Id != "-1" && game.Team2Id != "-1" {
					games = append(games, season.LeagueRounds[league][i].Games[j])
				}
			}
		}
	}
	return games
}

func isEarlyGame(hour, minute int) bool {
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
