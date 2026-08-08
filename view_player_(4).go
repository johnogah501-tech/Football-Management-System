package main

import "fmt"

// viewPlayers displays all players.

func viewPlayers(players []Player) {

	if len(players) == 0 {
		fmt.Println("No Players Found")
		fmt.Println()
		return
	}

	printPlayersTable(players)

}

// printPlayersTable displays players in a formatted table.

func printPlayersTable(players []Player) {

	border := "+------+--------------------+----------------------+--------------------+----------+------------+--------+----------+"

	fmt.Println(border)
	fmt.Printf("%65s\n", "PLAYERS TABLE")
	fmt.Println(border)

	fmt.Printf("| %-4s | %-18s | %-20s | %-18s | %-8s | %-10s | %-6s | %-8s | \n",
		"#",
		"Name",
		"Club",
		"Position",
		"Goals",
		"Assists",
		"Age",
		"Jersey",
	)
	fmt.Println(border)

	for i, player := range players {

		fmt.Printf(
			"| %-4d | %-18s | %-20s | %-18s | %-8d | %-10d | %-6d | %-8d |\n",

			i+1,
			player.Name,
			player.Club,
			player.Position,
			player.Goals,
			player.Assists,
			player.Age,
			player.JerseyNumber,
		)

	}

	fmt.Println(border)
	fmt.Println()
}
