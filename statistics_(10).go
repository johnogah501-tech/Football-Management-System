package main

import "fmt"

//for club statistics

func clubStatistics(players []Player) (int, int, int, float64) {

	var totalPlayers int
	var totalGoals int
	var totalAssists int
	var totalAge int

	for _, player := range players {

		totalPlayers++
		totalGoals += player.Goals
		totalAssists += player.Assists
		totalAge += player.Age

	}
	averageAge := float64(totalAge) / float64(totalPlayers)

	return totalPlayers, totalGoals, totalAssists, averageAge
}

func printClubStatistics(totalPlayers int, totalGoals int, totalAssists int, averageAge float64) {
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("         CLUB STATISTICS")
	fmt.Println("================================")
	fmt.Printf("Total Players    : %d\n", totalPlayers)
	fmt.Printf("Total Goals      : %d\n", totalGoals)
	fmt.Printf("Total Assists    : %d\n", totalAssists)
	fmt.Printf("Average Age      : %.2f\n", averageAge)
	fmt.Println("=================================")
	fmt.Println()
}

func viewClubStatistics(players []Player) {

	if len(players) == 0 {
		fmt.Println("No player available")
		fmt.Println()
		return
	}

	totalPlayers, totalGoals, totalAssists, averageAge := clubStatistics(players)

	printClubStatistics(
		totalPlayers,
		totalGoals,
		totalAssists,
		averageAge,
	)
}
