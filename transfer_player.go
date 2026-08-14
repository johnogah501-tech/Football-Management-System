package main

import "fmt"

// for transferring players

func inputNewClub() string {
	var club string

	fmt.Print("Enter new club : ")
	fmt.Scan(&club)

	return club
}

func transferClub(player *Player, newClub string) {
	player.Club = newClub
}

func transferPlayer(players []Player) {

	jerseyNumber := inputJerseyNumber()

	player := findPlayerByJerseyNumber(players, jerseyNumber)

	if player == nil {
		fmt.Println("Player not found")
		fmt.Println()
		return
	}

	newClub := inputNewClub()

	if player.Club == newClub {
		fmt.Println("Cannot transfer to the same club")
		fmt.Println()
		return
	}

	transferClub(player, newClub)

	fmt.Println("Player transfered successfully")
	fmt.Println()

	printPlayersTable([]Player{*player})
}
