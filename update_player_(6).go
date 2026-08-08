package main

import "fmt"

// Goal update functions.

func addGoals(player *Player, goals int) {
	player.Goals += goals
}

func inputGoalsScored() int {
	var goals int

	fmt.Print("Enter goals scored : ")
	fmt.Scan(&goals)

	return goals
}

func updatePlayerGoals(players []Player) {

	jerseyNumber := inputJerseyNumber()

	player := findPlayerByJerseyNumber(players, jerseyNumber)

	if player == nil {
		fmt.Println("Player not found")
		fmt.Println()
		return
	}

	goals := inputGoalsScored()

	addGoals(player, goals)

	fmt.Println("Goals added successfully")
	fmt.Println()

	printPlayersTable([]Player{*player})

}

// Assist update functions.

func inputPlayerAssists() int {
	var assists int

	fmt.Print("Enter player assists : ")
	fmt.Scan(&assists)

	return assists
}

func addAssists(player *Player, assists int) {
	player.Assists += assists
}

func updatePlayerAssists(players []Player) {

	jerseyNumber := inputJerseyNumber()

	player := findPlayerByJerseyNumber(players, jerseyNumber)

	if player == nil {
		fmt.Println("Player not found")
		fmt.Println()
		return
	}

	assists := inputPlayerAssists()

	addAssists(player, assists)

	fmt.Println("Assists added successfully")
	fmt.Println()

	printPlayersTable([]Player{*player})
}
