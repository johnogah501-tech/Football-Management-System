package main

import "fmt"

// inputPlayerDetails collects information for a new player.

func inputPlayerDetails(player *Player) {

	fmt.Println("=================================")
	fmt.Println("      PLAYER INFORMATION")
	fmt.Println("=================================")
	fmt.Print("Enter Name : ")
	fmt.Scan(&player.Name)

	fmt.Print("Enter Age : ")
	fmt.Scan(&player.Age)

	fmt.Print("Enter Jersey Number : ")
	fmt.Scan(&player.JerseyNumber)

	fmt.Print("Enter Position : ")
	fmt.Scan(&player.Position)

	fmt.Print("Enter Club : ")
	fmt.Scan(&player.Club)
	fmt.Println()
}

func jerseyNumberExists(players []Player, jerseyNumber int) bool {

	for _, player := range players {
		if player.JerseyNumber == jerseyNumber {
			return true
		}
	}
	return false
}

// add the new player

func addNewPlayer(players []Player) []Player {
	var player Player

	inputPlayerDetails(&player)

	if jerseyNumberExists(players, player.JerseyNumber) {
		fmt.Println("Jersey number already taken")
		fmt.Println()
		return players
	}

	player.Goals = 0

	player.Assists = 0

	players = append(players, player)

	printPlayerAdded(player)

	return players
}

func printPlayerAdded(player Player) {
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("    PLAYER SUCCESSFULLY ADDED")
	fmt.Println("=================================")
	fmt.Printf("Player Name      : %s\n", player.Name)
	fmt.Printf("Player Age       : %d\n", player.Age)
	fmt.Printf("Jersey Number    : %d\n", player.JerseyNumber)
	fmt.Printf("Player Position  : %s\n", player.Position)
	fmt.Printf("Player Club      : %s\n", player.Club)
	fmt.Printf("Goals Scored     : %d\n", player.Goals)
	fmt.Printf("Total Assists    : %d\n", player.Assists)
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("Welcome to the team!")
	fmt.Println()
}
