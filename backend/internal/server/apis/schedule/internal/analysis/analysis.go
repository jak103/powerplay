package analysis

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"math"
	"time"
)

func RunTimeAnalysis(games []models.Game) (structures.SeasonStats, map[string]structures.TeamStats) {
	seasonStats := structures.SeasonStats{
		TotalGames: len(games),
	}
	teamStats := make(map[string]structures.TeamStats)

	for _, game := range games {
		var team1Stats structures.TeamStats
		var team2Stats structures.TeamStats
		var ok bool

		if team1Stats, ok = teamStats[game.HomeTeam.Name]; !ok {
			team1Stats = newStats(game.HomeTeam.League.Name, game.HomeTeam.Name)
		}

		if team2Stats, ok = teamStats[game.AwayTeam.Name]; !ok {
			team2Stats = newStats(game.AwayTeam.League.Name, game.AwayTeam.Name)
		}

		earlyLateGames(game, &seasonStats, &team1Stats, &team2Stats)

		daysOfTheWeek(game, &team1Stats, &team2Stats)

		teamStats[game.HomeTeam.Name] = team1Stats
		teamStats[game.AwayTeam.Name] = team2Stats
	}

	seasonEarlyHigh := int(seasonStats.EarlyPercentage()*10.0) + 1 // TODO took a shortcut here and just hardcoded the 10 games

	for _, team := range teamStats {
		team.Balanced = team.EarlyGames <= seasonEarlyHigh
		teamStats[team.Name] = team
	}

	timeBetweenGames(games, teamStats)

	scoreAll(&seasonStats, teamStats)

	return seasonStats, teamStats
}

func Serialize(ts map[string]structures.TeamStats) []structures.TeamStats {
	var stats []structures.TeamStats
	for _, v := range ts {
		td := structures.TeamStats{
			Name:                    v.Name,
			League:                  v.League,
			EarlyGames:              v.EarlyGames,
			LateGames:               v.LateGames,
			DaysOfTheWeek:           v.DaysOfTheWeek,
			DaysBetweenGames:        v.DaysBetweenGames,
			AverageDaysBetweenGames: v.AverageDaysBetweenGames,
			Balanced:                v.Balanced,
		}
		stats = append(stats, td)
	}
	return stats
}

func newStats(league, team string) structures.TeamStats {
	return structures.TeamStats{
		League:        league,
		Name:          team,
		DaysOfTheWeek: make(map[time.Weekday]int),
	}
}

func earlyLateGames(game models.Game, season *structures.SeasonStats, team1, team2 *structures.TeamStats) {
	if IsEarlyGame(game.Start.Hour(), game.Start.Minute()) {
		team1.EarlyGames += 1
		team2.EarlyGames += 1
		season.EarlyGames += 1
	} else {
		team1.LateGames += 1
		team2.LateGames += 1
		season.LateGames += 1
	}
}

func daysOfTheWeek(game models.Game, team1, team2 *structures.TeamStats) {
	team1.DaysOfTheWeek[game.Start.Weekday()] += 1
	team2.DaysOfTheWeek[game.Start.Weekday()] += 1
}

func timeBetweenGames(games []models.Game, teamStats map[string]structures.TeamStats) {
	for team := range teamStats {
		stats := teamStats[team]
		for i := 1; i < len(games); i += 1 {
			previousGame := games[i-1]
			currentGame := games[i]

			betweenDuration := currentGame.Start.Sub(previousGame.Start)

			days := int(betweenDuration.Hours() / 24)
			stats.DaysBetweenGames = append(stats.DaysBetweenGames, days)

			stats.AverageDaysBetweenGames += float32(days)
		}

		stats.AverageDaysBetweenGames /= float32(len(stats.DaysBetweenGames))

		teamStats[team] = stats
	}
}

func IsEarlyGame(hour, minute int) bool {
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

func scoreAll(ss *structures.SeasonStats, ts map[string]structures.TeamStats) {
	// use the coefficient of variation so we can compare the variation across different pieces of data
	// see https://en.wikipedia.org/wiki/Coefficient_of_variation

	var earlyGamesCoefficientOfVariation float64
	var daysBetweenGamesMeanCoefficientOfVariation float64

	// compute the variation for early-late games for all teams over the season
	earlyGameCounts := make([]float64, len(ts))
	i := 0
	for _, teamStats := range ts {
		earlyGameCounts[i] = float64(teamStats.EarlyGames)
		i++
	}
	earlyGamesCoefficientOfVariation = coefficientOfVariation(earlyGameCounts)

	// compute the mean variation for time between games
	dayCoefficientsOfVariation := make([]float64, len(ts))
	i = 0
	for _, teamStats := range ts {
		daysBetweenGames := make([]float64, len(teamStats.DaysBetweenGames))
		for i, days := range teamStats.DaysBetweenGames {
			daysBetweenGames[i] = float64(days)
		}

		coefficient := coefficientOfVariation(daysBetweenGames)
		dayCoefficientsOfVariation[i] = coefficient
		i++
	}
	daysBetweenGamesMeanCoefficientOfVariation = mean(dayCoefficientsOfVariation)

	// TODO: maybe weight these
	ss.Score = earlyGamesCoefficientOfVariation + daysBetweenGamesMeanCoefficientOfVariation
}

func coefficientOfVariation(numbers []float64) float64 {
	mean := mean(numbers)
	return stddev(numbers, mean) / mean
}

func mean(numbers []float64) float64 {
	sum := 0.0
	for _, value := range numbers {
		sum += float64(value)
	}

	return sum / float64(len(numbers))
}

func stddev(numbers []float64, mean float64) float64 {
	sum := 0.0
	for _, value := range numbers {
		diff := float64(value) - mean
		sum += (diff * diff)
	}

	return math.Sqrt(sum / float64(len(numbers)))
}
