package main

import "fmt"

func main() {
	var budget, priceForOneNight float64
	var countOfNights, percentAddiitionalExpenses int
	fmt.Scanln(&budget)
	fmt.Scanln(&countOfNights)
	fmt.Scanln(&priceForOneNight)
	fmt.Scanln(&percentAddiitionalExpenses)

	// Calculations
	if countOfNights > 7 {
		priceForOneNight *= 0.95
	}

	priceForAllNights := float64(countOfNights) * priceForOneNight
	addiitionalExpenses := budget * float64(percentAddiitionalExpenses) / 100
	allExpenses := priceForAllNights + addiitionalExpenses

	// Ouput
	if allExpenses <= budget {
		leftMoney := budget - allExpenses
		fmt.Printf("Ivanovi will be left with %.2f leva after vacation.", leftMoney)
	} else {
		neededMoney := allExpenses - budget
		fmt.Printf("%.2f leva needed.", neededMoney)
	}

}
