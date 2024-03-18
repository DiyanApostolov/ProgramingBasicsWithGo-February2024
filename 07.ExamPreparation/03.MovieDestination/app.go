package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, priceOfDestination float64
	var destination, season string
	var countOfDays int
	fmt.Scanln(&budget)
	fmt.Scanln(&destination)
	fmt.Scanln(&season)
	fmt.Scanln(&countOfDays)

	//Calculations
	if season == "Summer" {
		switch destination {
		case "Dubai":
			priceOfDestination = 40000 * 0.7
		case "Sofia":
			priceOfDestination = 12500 * 1.25
		case "London":
			priceOfDestination = 20250
		}
	} else if season == "Winter" {
		switch destination {
		case "Dubai":
			priceOfDestination = 45000 * 0.7
		case "Sofia":
			priceOfDestination = 17000 * 1.25
		case "London":
			priceOfDestination = 24000
		}
	}

	finalPrice := priceOfDestination * float64(countOfDays)

	//Output
	diff := math.Abs(finalPrice - budget)

	if finalPrice <= budget {
		fmt.Printf("The budget for the movie is enough! We have %.2f leva left!", diff)
	} else {
		fmt.Printf("The director needs %.2f leva more!", diff)
	}
}
