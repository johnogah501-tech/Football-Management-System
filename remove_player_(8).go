package main

import "fmt"

// Player removal functions.

func findPlayerIndexByJerseyNumber(players []Player, jerseyNumber int) int {

	for i, player := range players {
		if player.JerseyNumber == jerseyNumber {
			return i
		}
	}
	return -1

}

func removePlayer(players []Player) []Player {

	jerseyNumber := inputJerseyNumber()

	player := findPlayerByJerseyNumber(players, jerseyNumber)

	if player == nil {
		fmt.Println("Player not found")
		fmt.Println()
		return players
	}

	// Make a copy BEFORE deleting
	removedPlayer := *player

	index := findPlayerIndexByJerseyNumber(players, jerseyNumber)

	players = append(players[:index], players[index+1:]...)

	fmt.Printf("%s removed successfully.\n", removedPlayer.Name)
	fmt.Println()

	return players
}
