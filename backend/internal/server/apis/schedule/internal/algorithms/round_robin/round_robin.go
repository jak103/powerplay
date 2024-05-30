package round_robin

import (
	"errors"
	"strconv"
	"time"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/optimize"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/utils/log"
)

func RoundRobin(leagues []models.League, iceTimes []string, numberOfGamesPerTeam int) ([]structures.Game, error) {
	if len(leagues) == 0 {
		return nil, errors.New("no leagues to generate games for")
	}
	if len(iceTimes) == 0 {
		return nil, errors.New("no ice times to assign")
	}
	if numberOfGamesPerTeam <= 0 {
		return nil, errors.New("no games per team to generate")
	}

	season, err := generateGames(leagues, numberOfGamesPerTeam)
	if err != nil {
		return nil, err
	}

	games, err := assignTimes(iceTimes, season, numberOfGamesPerTeam)
	if err != nil {
		return nil, err
	}

	optimizeSchedule(games)

	return games, nil
}

func optimizeSchedule(games []structures.Game) {
	if len(games) == 0 {
		log.Info("No games to optimize")
		return
	}
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

func generateGames(leagues []models.League, numberOfGamesPerTeam int) (structures.Season, error) {
	if len(leagues) == 0 {
		return structures.Season{}, errors.New("no leagues to generate games for")
	}
	season := structures.Season{LeagueRounds: make(map[string][]structures.Round)}

	for _, league := range leagues {
		numTeams := len(league.Teams)

		// Figure out how many rounds we need to run to get each team the number of games per season
		numberOfGamesPerTeam += ((numTeams * numberOfGamesPerTeam) - (numTeams/2)*(2*numberOfGamesPerTeam)) / 2

		log.Info("League %v games per round: %v\n", league.Name, numberOfGamesPerTeam)

		if numTeams%2 == 1 {
			league.Teams = append(league.Teams, models.Team{Name: "Bye", CorrelationId: "-1"})
			numTeams = len(league.Teams)
		}

		numberOfRounds := numberOfGamesPerTeam

		rounds := make([]structures.Round, numberOfRounds)

		for round := 0; round < numberOfRounds; round++ {
			rounds[round].Games = make([]structures.Game, numTeams/2)
			for i := 0; i < numTeams/2; i++ {
				team1 := league.Teams[i].DbModel.ID
				team1Name := league.Teams[i].Name
				team2 := league.Teams[numTeams-1-i].DbModel.ID
				team2Name := league.Teams[numTeams-1-i].Name

				rounds[round].Games[i] = newGame(league.Name, strconv.FormatUint(uint64(team1), 10), team1Name, strconv.FormatUint(uint64(team2), 10), team2Name)
			}

			rotateTeams(&league)
		}
		season.LeagueRounds[league.Name] = rounds
	}

	return season, nil
}

func assignTimes(times []string, season structures.Season, numberOfGamesPerTeam int) ([]structures.Game, error) {
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

		games[i].IsEarly = isEarlyGame(games[i].Start.Hour(), games[i].Start.Minute())
	}

	return games, nil
}

func getBalanceCount(teamStats *map[string]structures.TeamStats) int {
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

func newGame(league, team1, team1Name, team2, team2Name string) structures.Game {
	if team1 == "-1" || team2 == "-1" {
		return structures.Game{
			Team1Id:     team1,
			Team1Name:   team1Name,
			Team2Id:     team2,
			Team2Name:   team2Name,
			League:      league,
			Location:    "Bye",
			LocationUrl: "",
			EventType:   "Bye",
		}
	}
	return structures.Game{
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

func newGames(season *structures.Season, numberOfGamesPerTeam int) ([]structures.Game, error) {
	if season == nil {
		return nil, errors.New("no season to get games from")
	}
	if season.LeagueRounds == nil || len(season.LeagueRounds) == 0 {
		return nil, errors.New("no rounds to get games from")
	}
	games := make([]structures.Game, 0)
	for i := 0; i < numberOfGamesPerTeam; i += 1 { // Rounds // TODO This currently won't work if the leagues don't all have the same number of teams, fix this when needed (Balance by calculating the rate at which games have to be assigned, e.g. the average time between games to complete in the season from the number of first to last dates )
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
	return games, nil
}

func isEarlyGame(hour, minute int) bool {
	if hour < 20 {
		return true
	}
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
