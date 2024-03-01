package main

import (
	"fmt"
	"math"
)

func main() {
	const priceOfPuzzel float64 = 2.60
	const priceOfDoll float64 = 3
	const priceOfBear float64 = 4.10
	const priceOfMinion float64 = 8.20
	const priceOfTruck float64 = 2

	var priceOfExcursion float64
	var countOfPuzzels, countOfDolls, countOfBears, coutOfMinions, countOfTrucks int

	fmt.Scanln(&priceOfExcursion)
	fmt.Scanln(&countOfPuzzels)
	fmt.Scanln(&countOfDolls)
	fmt.Scanln(&countOfBears)
	fmt.Scanln(&coutOfMinions)
	fmt.Scanln(&countOfTrucks)

	sumPrices := float64(countOfPuzzels)*priceOfPuzzel + float64(countOfDolls)*priceOfDoll + float64(countOfBears)*priceOfBear + float64(coutOfMinions)*priceOfMinion + float64(countOfTrucks)*priceOfTruck

	countOfAlltoys := countOfPuzzels + countOfDolls + countOfBears + coutOfMinions + countOfTrucks

	if countOfAlltoys >= 50 {
		//sumPrices = sumPrices - (sumPrices * 0.25)
		//sumPrices = sumPrices * 0.75
		sumPrices *= 0.75
	}

	rent := sumPrices * 0.1

	finalIncom := sumPrices - rent
	diff := math.Abs(finalIncom - priceOfExcursion)

	if finalIncom >= priceOfExcursion {
		//moneyLeft := finalIncom - priceOfExcursion

		fmt.Printf("Yes! %.2f lv left.", diff)
		fmt.Printf("Yes! %.2f lv left.", diff)
	} else {
		//moneyNeeded := priceOfExcursion - finalIncom

		fmt.Printf("Not enough money! %.2f lv needed.", diff)
	}
}
