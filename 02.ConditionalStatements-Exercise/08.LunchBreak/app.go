package main

import (
	"fmt"
	"math"
)

func main() {
	var nameOfTvSeries string
	var timeOfTvSeries, timeOfLunchBreak int

	fmt.Scanln(&nameOfTvSeries)
	fmt.Scanln(&timeOfTvSeries)
	fmt.Scanln(&timeOfLunchBreak)

	lunchTime := float64(timeOfLunchBreak) / 8.0
	restTime := float64(timeOfLunchBreak) / 4.0

	allTime := float64(timeOfTvSeries) + lunchTime + restTime

	if allTime <= float64(timeOfLunchBreak) {
		timeLeft := math.Ceil(float64(timeOfLunchBreak) - allTime)

		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", nameOfTvSeries, timeLeft)
	} else {
		neededTime := math.Ceil(allTime - float64(timeOfLunchBreak))

		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", nameOfTvSeries, neededTime)
	}

}
