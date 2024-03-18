package main

import (
	"fmt"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	var tournamentDays, winCounter, loseCounter, winDayCounter, loseDayCounter int
	var sportName, winOrLose string
	var myMoney, currentDayMoney float64
	fmt.Scanln(&tournamentDays)

	for i := 0; i < tournamentDays; i++ {
		fmt.Scanln(&sportName)
		// scanner.Scan()
		// sportName = scanner.Text()

		for sportName != "Finish" {
			fmt.Scanln(&winOrLose)

			if winOrLose == "win" {
				winCounter++
				currentDayMoney += 20
			} else if winOrLose == "lose" {
				loseCounter++
			}

			fmt.Scanln(&sportName)
			// scanner.Scan()
			// sportName = scanner.Text()
		}

		if winCounter > loseCounter {
			currentDayMoney *= 1.1
			winDayCounter++
		} else {
			loseDayCounter++
		}

		myMoney += currentDayMoney
		winCounter = 0
		loseCounter = 0
		currentDayMoney = 0
	}

	if winDayCounter > loseDayCounter {
		myMoney *= 1.2
		fmt.Printf("You won the tournament! Total raised money: %.2f", myMoney)
	} else {
		fmt.Printf("You lost the tournament! Total raised money: %.2f", myMoney)
	}

}
