package main

import "fmt"

// Top scorer functions.

func findTopScorer(players []Player) Player {

	topScorer := players[0]

	for _, player := range players {

		if player.Goals > topScorer.Goals {
			topScorer = player
		}

	}

	return topScorer
}

func viewTopScorer(players []Player) {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return
	}

	topScorer := findTopScorer(players)

	printPlayersTable([]Player{topScorer})

}

// Top assists functions.

func findTopAssists(players []Player) Player {

	topAssists := players[0]

	for _, player := range players {

		if player.Assists > topAssists.Assists {
			topAssists = player
		}

	}

	return topAssists
}

func viewTopAssists(players []Player) {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return
	}

	topAssists := findTopAssists(players)

	printPlayersTable([]Player{topAssists})

}

// Top assists functions.

func findOldestPlayer(players []Player) Player {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return Player{}
	}

	oldestPlayer := players[0]

	for _, player := range players {

		if player.Age > oldestPlayer.Age {
			oldestPlayer = player
		}
	}

	return oldestPlayer
}

func viewOldestPlayer(players []Player) {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return
	}

	oldestPlayer := findOldestPlayer(players)

	printPlayersTable([]Player{oldestPlayer})
}

// Youngest player functions.

func findYoungestPlayer(players []Player) Player {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return Player{}
	}

	youngestPlayer := players[0]

	for _, player := range players {

		if player.Age < youngestPlayer.Age {
			youngestPlayer = player
		}
	}

	return youngestPlayer
}

func viewYoungestPlayer(players []Player) {

	if len(players) == 0 {
		fmt.Println("No players available.")
		fmt.Println()
		return
	}

	youngestPlayer := findYoungestPlayer(players)

	printPlayersTable([]Player{youngestPlayer})
}
