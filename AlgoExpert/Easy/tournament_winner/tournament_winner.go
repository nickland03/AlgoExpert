package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	currentBestTeam := ""
	teams := map[string]int{currentBestTeam: 0}

	for i := 0; i < len(results); i++ {
		homeTeam := competitions[i][0]
		awayTeam := competitions[i][1]
		var winner string

		if results[i] == HOME_TEAM_WON {
			winner = homeTeam
		} else {
			winner = awayTeam
		}

		teams[winner] += 3

		if teams[winner] > teams[currentBestTeam] {
			currentBestTeam = winner
		}
	}

	return currentBestTeam
}