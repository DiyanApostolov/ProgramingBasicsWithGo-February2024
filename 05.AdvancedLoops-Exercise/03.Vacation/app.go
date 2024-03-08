package main

import (
	"fmt"
)

func main() {
	var moneyForTrip, savedMoney, currentMoney float64
	var move string
	var spendCounter, allDayCounter int

	fmt.Scanln(&moneyForTrip)
	fmt.Scanln(&savedMoney)

	for savedMoney < moneyForTrip {
		fmt.Scanln(&move)
		fmt.Scanln(&currentMoney)
		allDayCounter++

		if move == "spend" {
			spendCounter++
			savedMoney -= currentMoney
			if savedMoney < 0 {
				savedMoney = 0
			}

			if spendCounter == 5 {
				fmt.Println("You can't save the money.")
				fmt.Println(allDayCounter)
				break
			}
		} else if move == "save" {
			spendCounter = 0
			savedMoney += currentMoney
		}

	}
	
	if savedMoney >= moneyForTrip {
		fmt.Printf("You saved the money for %d days.", allDayCounter)
	}
	

}
