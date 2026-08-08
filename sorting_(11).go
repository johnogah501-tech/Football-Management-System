package main

import "fmt"

// Player sorting functions.

func sortPlayersByGoals(players []Player) {

	for i := range players {

		for j := i + 1; j < len(players); j++ {

			// Swap the two players.
			if players[i].Goals < players[j].Goals {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
}

func viewPlayersSortedByGoals(players []Player) {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return
	}

	sortedPlayers := append([]Player{}, players...)

	sortPlayersByGoals(sortedPlayers)

	printPlayersTable(sortedPlayers)
}
