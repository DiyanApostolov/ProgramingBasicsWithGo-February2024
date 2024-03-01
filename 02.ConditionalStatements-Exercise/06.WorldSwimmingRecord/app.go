package main

import (
	"fmt"
	"math"
)

func main() {
	var recordInSeconds, distanceInMetres, secondsForOneMeter float64

	fmt.Scanln(&recordInSeconds)
	fmt.Scanln(&distanceInMetres)
	fmt.Scanln(&secondsForOneMeter)

	timeForSwiming := distanceInMetres * secondsForOneMeter
	delay := math.Floor(distanceInMetres / 15) * 12.5

	finalTime := timeForSwiming + delay

	if finalTime < recordInSeconds {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", finalTime)
	} else {
		neededTime := finalTime - recordInSeconds

		fmt.Printf("No, he failed! He was %.2f seconds slower.", neededTime)
	}
}
