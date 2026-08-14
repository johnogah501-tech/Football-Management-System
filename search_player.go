package main

import "fmt"

// inputJerseyNumber asks for a player's jersey number.

func inputJerseyNumber() int {
	var jerseyNumber int

	fmt.Print("Enter Player's Jersey Number : ")
	fmt.Scan(&jerseyNumber)

	return jerseyNumber
}

// findPlayerByJerseyNumber finds a player by jersey number.

func findPlayerByJerseyNumber(players []Player, jerseyNumber int) *Player {
	for i := range players {
		if players[i].JerseyNumber == jerseyNumber {
			return &players[i]
		}
	}
	return nil
}

// searchPlayer searches for and displays a player.

func searchPlayer(players []Player) {

	jerseyNumber := inputJerseyNumber()

	player := findPlayerByJerseyNumber(players, jerseyNumber)

	if player == nil {
		fmt.Println("Player not found")
		fmt.Println()
		return
	}

	printPlayersTable([]Player{*player})
}
