package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	var stepsToHome int
	steps := 10000

	for steps > 0 {
		fmt.Scanln(&input)

		if input == "GoingHome" {
			fmt.Scanln(&stepsToHome)
			steps -= stepsToHome

			if steps > 0 {
				fmt.Printf("%d more steps to reach goal.", steps)
				break
			}
		} else {
			currentSteps, _ := strconv.Atoi(input)
			steps -= currentSteps
		}
	}

	if steps <= 0 {
		fmt.Println("Goal reached! Good job!")
		fmt.Printf("%.0f steps over the goal!", math.Abs(float64(steps)))
	}

}
