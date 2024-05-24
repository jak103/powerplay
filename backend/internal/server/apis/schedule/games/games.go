package games

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/util"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/write"
	"time"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/optimize"
	"github.com/jak103/powerplay/internal/server/apis/schedule/helpers/read"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/games", auth.Authenticated, handleGames)
}

func handleGames(c *fiber.Ctx) error {
	log.Info("Reading body\n")
	// TODO numberOfGamesPerTeam is read from the request for now. We need to read this from the ice_times.csv file.
	seasonName, numberOfGamesPerTeam, err := readBody(c)
	if err != nil {
		log.Error("Error reading body: %v\n", err)
		return err
	}

	if seasonName == "" {
		log.Error("Season file name is empty\n")
		return responder.BadRequest(c, "Season file name is empty")
	}

	if numberOfGamesPerTeam == 0 {
		log.Error("Number of games per team is 0\n")
		return responder.BadRequest(c, "Number of games per team is 0")
	}

	log.Info("Reading config file for season\n")
	seasonConfig, err := read.SeasonConfig(seasonName)
	if err != nil {
		log.Error("Error reading file: %v\n", err)
		return responder.BadRequest(c, "Error reading file")
	}

	log.Info("Generating games\n")
	season, err := generateGames(seasonConfig.Leagues, numberOfGamesPerTeam)

	log.Info("Assigning ice times\n")
	games, err := assignTimes(seasonConfig.IceTimes, season, numberOfGamesPerTeam)
	if err != nil {
		log.Error("Error assigning ice times: %v\n", err)
		return responder.BadRequest(c, "Error assigning ice times")
	}

	log.Info("Optimizing schedule\n")
	optimizeSchedule(games, numberOfGamesPerTeam)

	log.Info("Writing csv\n")
	err = write.Csv(games, "schedule.csv")
	if err != nil {
		log.Error("Error writing csv: %v\n", err)
		return responder.InternalServerError(c)
	}

	//log.Info("Saving to database\n")
	//err = write.ToDb(c, games)
	//if err != nil {
	//	log.Error("Error saving to database: %v\n", err)
	//	return responder.InternalServerError(c)
	//}

	return responder.Ok(c, "Schedule generated at schedule.csv and saved to database")
}

func readBody(c *fiber.Ctx) (string, int, error) {
	type BodyDto struct {
		SeasonName           string `json:"seasonName"`
		NumberOfGamesPerTeam int    `json:"numberOfGamesPerTeam"`
	}
	body := c.Body()
	var bodyDto BodyDto
	err := json.Unmarshal(body, &bodyDto)
	if err != nil {
		return "", 0, responder.BadRequest(c, "Error reading body")
	}

	return bodyDto.SeasonName, bodyDto.NumberOfGamesPerTeam, nil
}

func assignLockerRooms(game *models.Game) {
	if game.IsEarly {
		game.Team1LockerRoom = "3" // Home team
		game.Team2LockerRoom = "1" // Away team
	} else {
		game.Team1LockerRoom = "5"
		game.Team2LockerRoom = "2"
	}
}

func optimizeSchedule(games []models.Game, numberOfGamesPerTeam int) {
	if len(games) == 0 {
		log.Info("No games to optimize")
		return
	}
	log.Info("Pre-optimization analysis")
	seasonStats, teamStats := analysis.RunTimeAnalysis(games, numberOfGamesPerTeam)

	// Need to make sure games are balanced in
	// - Early / late
	// - Days between games
	balanceCount := getBalanceCount(&teamStats)
	lastBalanceCount := -1

	for count := 0; balanceCount != lastBalanceCount && count < 25; count++ {
		optimize.Schedule(games, seasonStats, teamStats)

		log.Info("Post-optimization analysis")
		seasonStats, teamStats = analysis.RunTimeAnalysis(games, numberOfGamesPerTeam)

		lastBalanceCount = balanceCount
		balanceCount := getBalanceCount(&teamStats)

		log.Info("Balanced count: %v\n", balanceCount)
	}
}

func generateGames(leagues []models.League, numberOfGamesPerTeam int) (models.Season, error) {
	if len(leagues) == 0 {
		return models.Season{}, errors.New("no leagues to games games for")
	}
	season := models.Season{LeagueRounds: make(map[string][]models.Round)}

	for _, league := range leagues {
		numTeams := len(league.Teams)

		// Figure out how many rounds we need to run to get each team the number of games per season
		numberOfGamesPerTeam += ((numTeams * numberOfGamesPerTeam) - (numTeams/2)*(2*numberOfGamesPerTeam)) / 2

		log.Info("League %v games per round: %v\n", league.Name, numberOfGamesPerTeam)

		if numTeams%2 == 1 {
			league.Teams = append(league.Teams, models.Team{IsBye: true})
			numTeams = len(league.Teams)
		}

		numberOfRounds := numberOfGamesPerTeam

		rounds := make([]models.Round, numberOfRounds)

		for round := 0; round < numberOfRounds; round++ {
			rounds[round].Games = make([]models.Game, numTeams/2)
			for i := 0; i < numTeams/2; i++ {
				team1 := league.Teams[i]
				team2 := league.Teams[numTeams-1-i]

				rounds[round].Games[i] = newGame(league.Name, team1, team2)
			}

			rotateTeams(&league)
		}
		season.LeagueRounds[league.Name] = rounds
	}

	return season, nil
}

func assignTimes(times []string, season models.Season, numberOfGamesPerTeam int) ([]models.Game, error) {
	if len(times) == 0 {
		return nil, errors.New("no times to assign")
	}
	if season.LeagueRounds == nil {
		return nil, errors.New("no games to assign times to")
	}
	if len(times) < numberOfGamesPerTeam {
		return nil, errors.New("not enough times to assign")
	}

	games, err := newGames(&season, numberOfGamesPerTeam)
	if err != nil {
		return nil, err
	}

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

		games[i].IsEarly = util.IsEarlyGame(games[i].Start.Hour(), games[i].Start.Minute())
	}

	return games, nil
}

func getBalanceCount(teamStats *map[string]models.TeamStats) int {
	if teamStats == nil {
		return 0
	}
	balanceCount := 0
	for _, team := range *teamStats {
		if team.Balanced {
			balanceCount++
		}
	}
	return balanceCount
}

func rotateTeams(league *models.League) {
	if len(league.Teams) <= 2 {
		return
	}
	// Rotate teams except the first one
	lastTeam := league.Teams[len(league.Teams)-1]
	copy(league.Teams[2:], league.Teams[1:len(league.Teams)-1])
	league.Teams[1] = lastTeam
}

func newGame(league string, team1 models.Team, team2 models.Team) models.Game {
	if team1.IsBye || team2.IsBye {
		return models.Game{
			Teams:  []models.Team{team1, team2},
			League: league,
			IsBye:  true,
		}
	}
	game := models.Game{
		Teams:  []models.Team{team1, team2},
		League: league,
	}
	assignLockerRooms(&game)
	return game
}

func newGames(season *models.Season, numberOfGamesPerTeam int) ([]models.Game, error) {
	if season == nil {
		return nil, errors.New("no season to get games from")
	}
	if season.LeagueRounds == nil || len(season.LeagueRounds) == 0 {
		return nil, errors.New("no rounds to get games from")
	}
	games := make([]models.Game, 0)
	for i := 0; i < numberOfGamesPerTeam; i += 1 { // Rounds // TODO This currently won't work if the leagues don't all have the same number of teams, fix this when needed (Balance by calculating the rate at which games have to be assigned, e.g. the average time between games to complete in the season from the number of first to last dates )
		for _, league := range []string{"A", "C", "B", "D"} { // Alternate leagues, so if you play in two leagues, you don't play back to back
			if season.LeagueRounds[league] == nil || len(season.LeagueRounds[league]) <= i {
				continue
			}
			for j, game := range season.LeagueRounds[league][i].Games {
				if !game.Teams[0].IsBye && !game.Teams[1].IsBye {
					games = append(games, season.LeagueRounds[league][i].Games[j])
				}
			}
		}
	}
	return games, nil
}
