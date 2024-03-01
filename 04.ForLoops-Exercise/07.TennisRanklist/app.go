package main

import "fmt"

func main() {
	var countOfTournaments, startingPoints, tournametPoint, winTournaments int
	var result string

	fmt.Scanln(&countOfTournaments)
	fmt.Scanln(&startingPoints)

	for i := 0; i < countOfTournaments; i++ {
		fmt.Scanln(&result)

		switch result {
		case "W":
			tournametPoint += 2000
			winTournaments++
		case "F":
			tournametPoint += 1200
		case "SF":
			tournametPoint += 720
		}
	}

	averagePoints := tournametPoint / countOfTournaments

	var percentWinTournamets float64 = float64(winTournaments) / float64(countOfTournaments) * 100
	var finalPoints = startingPoints + tournametPoint

	fmt.Printf("Final points: %d\n", finalPoints)
	fmt.Printf("Average points: %d\n", averagePoints)
	fmt.Printf("%.2f%%", percentWinTournamets)

}
