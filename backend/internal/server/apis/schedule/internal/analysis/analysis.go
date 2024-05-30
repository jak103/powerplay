package analysis

import (
	"fmt"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/utils/log"
	"sort"
	"time"
)

func RunTimeAnalysis(games []structures.Game) (structures.SeasonStats, map[string]structures.TeamStats) {
	seasonStats := structures.SeasonStats{
		TotalGames: len(games),
	}
	teamStats := make(map[string]structures.TeamStats)

	for _, game := range games {
		var team1Stats structures.TeamStats
		var team2Stats structures.TeamStats
		var ok bool

		if team1Stats, ok = teamStats[game.Team1Name]; !ok {
			team1Stats = newStats(game.League, game.Team1Name)
		}

		if team2Stats, ok = teamStats[game.Team2Name]; !ok {
			team2Stats = newStats(game.League, game.Team2Name)
		}

		team1Stats.Games = append(team1Stats.Games, game)
		team2Stats.Games = append(team2Stats.Games, game)

		earlyLateGames(game, &seasonStats, &team1Stats, &team2Stats)

		daysOfTheWeek(game, &team1Stats, &team2Stats)

		teamStats[game.Team1Name] = team1Stats
		teamStats[game.Team2Name] = team2Stats
	}

	seasonEarlyHigh := int(seasonStats.EarlyPercentage()*10.0) + 1 // TODO took a shortcut here and just hardcoded the 10 games

	for _, team := range teamStats {
		team.Balanced = team.EarlyGames <= seasonEarlyHigh
		teamStats[team.Name] = team
	}

	timeBetweenGames(teamStats)

	printStats(seasonStats, teamStats)

	return seasonStats, teamStats
}

func newStats(league, team string) structures.TeamStats {
	return structures.TeamStats{
		League:        league,
		Name:          team,
		DaysOfTheWeek: make(map[time.Weekday]int),
		Games:         make([]structures.Game, 0),
	}
}

func earlyLateGames(game structures.Game, season *structures.SeasonStats, team1, team2 *structures.TeamStats) {
	if game.IsEarly {
		team1.EarlyGames += 1
		team2.EarlyGames += 1
		season.EarlyGames += 1
	} else {
		team1.LateGames += 1
		team2.LateGames += 1
		season.LateGames += 1
	}
}

func daysOfTheWeek(game structures.Game, team1, team2 *structures.TeamStats) {
	team1.DaysOfTheWeek[game.Start.Weekday()] += 1
	team2.DaysOfTheWeek[game.Start.Weekday()] += 1
}

func timeBetweenGames(teamStats map[string]structures.TeamStats) {
	for team := range teamStats {
		stats := teamStats[team]
		for i := 1; i < len(stats.Games); i += 1 {
			previousGame := stats.Games[i-1]
			currentGame := stats.Games[i]

			betweenDuration := currentGame.Start.Sub(previousGame.Start)

			days := int(betweenDuration.Hours() / 24)
			stats.DaysBetweenGames = append(stats.DaysBetweenGames, days)

			stats.AverageDaysBetweenGames += float32(days)
		}

		stats.AverageDaysBetweenGames /= float32(len(stats.DaysBetweenGames))

		teamStats[team] = stats
	}
}

func printStats(seasonStats structures.SeasonStats, teamStats map[string]structures.TeamStats) {
	log.Debug("Early games: %v/%v (%v%%)\n", seasonStats.EarlyGames, seasonStats.TotalGames, (float32(seasonStats.EarlyGames)/float32(seasonStats.TotalGames))*100)
	log.Debug("Late  games: %v/%v (%v%%)\n", seasonStats.LateGames, seasonStats.TotalGames, (float32(seasonStats.LateGames)/float32(seasonStats.TotalGames))*100)

	for _, league := range []string{"A", "B", "C", "D"} {
		log.Debug("%v league:\n", league)

		teams := make([]string, 0)
		for key := range teamStats {
			teams = append(teams, key)
		}
		sort.Strings(teams)

		for _, team := range teams {
			stats := teamStats[team]
			if stats.League == league {
				log.Debug("%v: (%v-%v)\n", team, stats.EarlyGames, stats.LateGames)

				log.Debug("Days of the week: ")
				for _, dotw := range []time.Weekday{1, 2, 3, 4, 5, 6} {
					day := fmt.Sprintf("%v", dotw)
					log.Debug("%c:%v ", day[0], stats.DaysOfTheWeek[dotw])
				}
				log.Debug("\n")

				log.Debug("Days between games: ")
				for i, daysBetween := range stats.DaysBetweenGames {
					if i < len(stats.DaysBetweenGames)-1 {
						log.Debug("%v-", daysBetween)
					} else {
						log.Debug("%v", daysBetween)
					}
				}
				log.Debug("\n")

				log.Debug("Average days between games: %0.1f\n", stats.AverageDaysBetweenGames)
			}
		}
	}
}
