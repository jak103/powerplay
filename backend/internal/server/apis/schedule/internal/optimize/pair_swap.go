package optimize

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/utils/log"
	"sort"
)

var oMap map[uint]bool

func PairOptimizeSchedule(games []models.Game) {
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

	for count := 0; balanceCount != lastBalanceCount && count < 3; count++ {
		pairSwapSchedule(games, seasonStats, teamStats)

		log.Info("Post-optimization analysis")
		seasonStats, teamStats = analysis.RunTimeAnalysis(games)

		lastBalanceCount = balanceCount
		balanceCount := getBalanceCount(&teamStats)

		log.Info("Balanced count: %v\n", balanceCount)
	}
}

func pairSwapSchedule(games []models.Game, seasonStats structures.SeasonStats, teamStats map[string]structures.TeamStats) {
	resetOptimized(games)
	log.Info("OPT: Early games percent: %v%%\n", seasonStats.EarlyPercentage()*100)
	seasonEarlyHigh := int(seasonStats.EarlyPercentage()*10.0) + 1 // TODO took a shortcut here and just hardcoded the 10 games
	seasonEarlyLow := int(seasonStats.EarlyPercentage() * 10.0)    // - 1

	// implemented as a series of constraints that have to pass in order to swap games

	log.Info("High count: %v\n", seasonEarlyHigh)
	log.Info("Low count: %v\n", seasonEarlyLow)

	teams := make([]string, 0)
	for key := range teamStats {
		teams = append(teams, key)
	}
	sort.Strings(teams)

	// Go through list of teams
	for _, team := range teams { // Sorted teams
		stats := teamStats[team]
		// for team, stats := range teamStats { // Unsorted teams

		// Does this team need to be balanced?
		needsBalance, tooManyEarly := needsToBeBalanced(stats, seasonEarlyHigh, seasonEarlyLow)
		if !needsBalance {
			log.Info("%s (%v-%v) is balanced\n", team, stats.EarlyGames, stats.LateGames)
			continue
		}

		log.Info("Need to balance: %v (%v-%v)\n", team, stats.EarlyGames, stats.LateGames)
		swaps := false
		// Look through all of the teams games and swaps
		for i := 1; i < len(games); i++ {
			if games[i].HomeTeam.Name == team || games[i].AwayTeam.Name == team {
				// find a good candidate to swap games with
				// Will it improve that balance
				if analysis.IsEarlyGame(games[i-1].Start.Hour(), games[i-1].Start.Minute()) == tooManyEarly {
					log.Info("Can't swap games because it won't improve balance\n")
					continue
				}

				// Does it force the swapped teams out of balance?
				if !correctBalanceDirection(teamStats, seasonEarlyHigh, seasonEarlyLow, games[i-1].HomeTeam.Name, games[i-1].AwayTeam.Name, tooManyEarly) {
					log.Info("Not swapping games because it won't help: %v (%v-%v) v (%v-%v)\n", games[i-1],
						teamStats[games[i-1].HomeTeam.Name].EarlyGames,
						teamStats[games[i-1].HomeTeam.Name].LateGames,
						teamStats[games[i-1].AwayTeam.Name].EarlyGames,
						teamStats[games[i-1].AwayTeam.Name].LateGames)
					continue
				}

				// Don't swap games that are already optimized
				if oMap[games[i].ID] && oMap[games[i-1].ID] {
					log.Info("Can't swap games because it has already been swapped\n")
					continue
				}

				// If we made it here, then all the constraints are met, go ahead and swap games
				// swap games
				log.Info("Swapping %v with %v\n", games[i], games[i-1])
				// log.Info("Do stats match? %v v %v\n", teamStats[games[i].Team1Name].Name, teamStats[games[i].Team2Name].Name)
				// log.Info("Do stats match? %v v %v\n", teamStats[games[i-1].Team1Name].Name, teamStats[games[i-1].Team2Name].Name)
				updateStats(teamStats, games, i, i-1)
				swapGames(games, i, i-1)
				swaps = true
			}
		}

		// Look through all of the teams games and swaps
		for i := 0; i < len(games)-1; i++ {
			if games[i].HomeTeam.Name == team || games[i].AwayTeam.Name == team {
				// find a good candidate to swap games with
				// Will it improve that balance
				if analysis.IsEarlyGame(games[i+1].Start.Hour(), games[i+1].Start.Minute()) == tooManyEarly {
					log.Info("Can't swap games because it won't improve balance\n")
					continue
				}

				// Does it force the swapped teams out of balance?
				if !correctBalanceDirection(teamStats, seasonEarlyHigh, seasonEarlyLow, games[i+1].HomeTeam.Name, games[i+1].AwayTeam.Name, tooManyEarly) {
					log.Info("Not swapping games because it won't help: %v (%v-%v) v (%v-%v)\n", games[i+1],
						teamStats[games[i+1].HomeTeam.Name].EarlyGames,
						teamStats[games[i+1].HomeTeam.Name].LateGames,
						teamStats[games[i+1].AwayTeam.Name].EarlyGames,
						teamStats[games[i+1].AwayTeam.Name].LateGames)
					continue
				}

				// Don't swap games that are already swapped
				if oMap[games[i].ID] && oMap[games[i+1].ID] {
					log.Info("Can't swap games because it has already been swapped\n")
					continue
				}

				// If we made it here, then all the constraints are met, go ahead and swap games
				// swap games
				log.Info("Swapping %v with %v\n", games[i], games[i+1])
				// log.Info("Do stats match? %v v %v\n", teamStats[games[i].Team1Name].Name, teamStats[games[i].Team2Name].Name)
				// log.Info("Do stats match? %v v %v\n", teamStats[games[i+1].Team1Name].Name, teamStats[games[i+1].Team2Name].Name)
				updateStats(teamStats, games, i, i+1)
				swapGames(games, i, i+1)
				swaps = true
			}
		}

		if !swaps {
			// Look through all of the teams games and swaps
			for i := 1; i < len(games); i++ {
				if games[i].HomeTeam.Name == team || games[i].AwayTeam.Name == team {
					if teamStats[team].EarlyGames > seasonEarlyHigh {
						log.Info("No swaps and we aren't balanced\n")
						// If we made it here, then all the constraints are met, go ahead and swap games
						// swap games
						log.Info("Swapping %v with %v\n", games[i], games[i-1])
						// log.Info("Do stats match? %v v %v\n", teamStats[games[i].Team1Name].Name, teamStats[games[i].Team2Name].Name)
						// log.Info("Do stats match? %v v %v\n", teamStats[games[i-1].Team1Name].Name, teamStats[games[i-1].Team2Name].Name)
						updateStats(teamStats, games, i, i-1)
						swapGames(games, i, i-1)
						break
					} else if teamStats[team].EarlyGames < seasonEarlyLow {
						log.Info("No swaps and we aren't balanced\n")
						// If we made it here, then all the constraints are met, go ahead and swap games
						// swap games
						log.Info("Swapping %v with %v\n", games[i], games[i-1])
						// log.Info("Do stats match? %v v %v\n", teamStats[games[i].Team1Name].Name, teamStats[games[i].Team2Name].Name)
						// log.Info("Do stats match? %v v %v\n", teamStats[games[i-1].Team1Name].Name, teamStats[games[i-1].Team2Name].Name)
						updateStats(teamStats, games, i-1, i)
						swapGames(games, i-1, i)
						break
					}
				}
			}
		}

		teamStats[team] = stats // update the stats
	}
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

func updateStats(teamStats map[string]structures.TeamStats, games []models.Game, i, j int) {
	game1Team1Stats := teamStats[games[i].HomeTeam.Name]
	game1Team2Stats := teamStats[games[i].AwayTeam.Name]
	game2Team1Stats := teamStats[games[j].HomeTeam.Name]
	game2Team2Stats := teamStats[games[j].AwayTeam.Name]

	if analysis.IsEarlyGame(games[i].Start.Hour(), games[i].Start.Minute()) {
		game1Team1Stats.EarlyGames--
		game1Team2Stats.EarlyGames--
		game1Team1Stats.LateGames++
		game1Team2Stats.LateGames++

		game2Team1Stats.EarlyGames++
		game2Team2Stats.EarlyGames++
		game2Team1Stats.LateGames--
		game2Team2Stats.LateGames--
	} else {
		game1Team1Stats.EarlyGames++
		game1Team2Stats.EarlyGames++
		game1Team1Stats.LateGames--
		game1Team2Stats.LateGames--

		game2Team1Stats.EarlyGames--
		game2Team2Stats.EarlyGames--
		game2Team1Stats.LateGames++
		game2Team2Stats.LateGames++
	}

	teamStats[games[i].HomeTeam.Name] = game1Team1Stats
	teamStats[games[i].AwayTeam.Name] = game1Team2Stats
	teamStats[games[j].HomeTeam.Name] = game2Team1Stats
	teamStats[games[j].AwayTeam.Name] = game2Team2Stats

	log.Info("Now %v (%v-%v) v %v (%v-%v)\n", game1Team1Stats.Name, game1Team1Stats.EarlyGames, game1Team1Stats.LateGames, game1Team2Stats.Name, game1Team2Stats.EarlyGames, game1Team2Stats.LateGames)
	log.Info("Now %v (%v-%v) v %v (%v-%v)\n", game2Team1Stats.Name, game2Team1Stats.EarlyGames, game2Team1Stats.LateGames, game2Team2Stats.Name, game2Team2Stats.EarlyGames, game2Team2Stats.LateGames)
}

func swapGames(games []models.Game, i, j int) {
	team1 := games[i].HomeTeam
	team2 := games[i].AwayTeam
	league := games[i].HomeTeam.League

	games[i].HomeTeam = games[j].HomeTeam
	games[i].AwayTeam = games[j].AwayTeam
	games[i].HomeTeam.League = games[j].HomeTeam.League

	games[j].HomeTeam = team1
	games[j].AwayTeam = team2
	games[j].HomeTeam.League = league

	// mark the swapped games as optimized so they can't be swapped again
	oMap[games[j].ID] = true
	oMap[games[i].ID] = true
}

func correctBalanceDirection(teamStats map[string]structures.TeamStats, seasonEarlyHigh, seasonEarlyLow int, team1, team2 string, tooManyEarly bool) bool {
	// tooManyEarly means the team trying to swap has too many early games
	team1Stats := teamStats[team1]
	team2Stats := teamStats[team2]
	// If this is going to move another team that needs to go early, do it
	if tooManyEarly && (teamStats[team1].EarlyGames < seasonEarlyLow || teamStats[team2].EarlyGames < seasonEarlyLow) {
		log.Info("Swap down with %s (%v-%v) v %s (%v-%v)\n", team1, team1Stats.EarlyGames, team1Stats.LateGames, team2, team2Stats.EarlyGames, team2Stats.LateGames)
		return true
	}

	// If this is going to move another games later, do it
	if !tooManyEarly && (teamStats[team1].EarlyGames > seasonEarlyHigh || teamStats[team2].EarlyGames > seasonEarlyHigh) {
		log.Info("Swap up with %s (%v-%v) v %s (%v-%v)\n", team1, team1Stats.EarlyGames, team1Stats.LateGames, team2, team2Stats.EarlyGames, team2Stats.LateGames)
		return true
	}

	// Otherwise, don't do it
	return false
}

func needsToBeBalanced(stats structures.TeamStats, seasonEarlyHigh, seasonEarlyLow int) (bool, bool) {
	if stats.EarlyGames > seasonEarlyHigh {
		return true, true // needs to be balanced, and has too many early games
	} else if stats.EarlyGames < seasonEarlyLow {
		return true, false // needs to be balanced and doen's have enough early games
	}

	// doesn't need to balance
	return false, false
}

func resetOptimized(games []models.Game) {
	oMap = make(map[uint]bool)
	for i := range games {
		oMap[games[i].ID] = false
	}
}
