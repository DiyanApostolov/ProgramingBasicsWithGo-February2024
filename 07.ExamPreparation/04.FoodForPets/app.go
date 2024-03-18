package main

import (
	"fmt"
	"math"
)

func main() {
	var countOfDays, dogFoodPerDay, catFoodPerDay, allDogFood, allCatFood int
	var foodAmount, biscuits float64
	fmt.Scanln(&countOfDays)
	fmt.Scanln(&foodAmount)

	for day := 1; day <= countOfDays; day++ {
		fmt.Scanln(&dogFoodPerDay)
		fmt.Scanln(&catFoodPerDay)
		allDogFood += dogFoodPerDay
		allCatFood += catFoodPerDay

		if day%3 == 0 {
			biscuits += (float64(dogFoodPerDay) + float64(catFoodPerDay))*0.1
		}
	}

	finalBiscuits := math.Round(biscuits)
	percentAllFoodEaten := (float64(allDogFood) + float64(allCatFood)) / foodAmount * 100
	percentDogEatenDoof := float64(allDogFood) / (float64(allDogFood) + float64(allCatFood)) * 100
	percentCatEatenDoof := float64(allCatFood) / (float64(allDogFood) + float64(allCatFood)) * 100

	// Output
	fmt.Printf("Total eaten biscuits: %.0fgr.\n", finalBiscuits)
	fmt.Printf("%.2f%% of the food has been eaten.\n", percentAllFoodEaten)
	fmt.Printf("%.2f%% eaten from the dog.\n", percentDogEatenDoof)
	fmt.Printf("%.2f%% eaten from the cat.", percentCatEatenDoof)
}
