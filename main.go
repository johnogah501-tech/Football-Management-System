//Finished Football Management system V1.0 on 01/Aug/2026.
//Learned about pointers and deleting slice)

//Seperated and Polished on 08/Aug/2026.

package main

import "fmt"

//for the heart of the program (main function)

func main() {

	players := []Player{}

	for {

		showMenu()

		var choice int

		fmt.Print("Choose an option : ")
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {

		case 1:
			players = addNewPlayer(players)

		case 2:
			viewPlayers(players)

		case 3:
			searchPlayer(players)

		case 4:
			updatePlayerGoals(players)

		case 5:
			updatePlayerAssists(players)

		case 6:
			transferPlayer(players)

		case 7:
			players = removePlayer(players)

		case 8:
			viewTopScorer(players)

		case 9:
			viewTopAssists(players)

		case 10:
			viewOldestPlayer(players)

		case 11:
			viewYoungestPlayer(players)

		case 12:
			viewClubStatistics(players)

		case 13:
			viewPlayersSortedByGoals(players)

		case 14:
			fmt.Println("Thank you for using Football Club Management System.")
			fmt.Println()
			return

		default:
			fmt.Println("========================================")
			fmt.Println("            INVALID OPTION")
			fmt.Println("========================================")
			fmt.Println("Please select an option between 1 and 14")
		}

	}

}
