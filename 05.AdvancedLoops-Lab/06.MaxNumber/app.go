package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	var maxNumber = math.MinInt32

	for input != "Stop" {
		currentNum, _ := strconv.Atoi(input)
		
		if currentNum > maxNumber{
			maxNumber = currentNum
		}

		fmt.Scanln(&input)
	}

	fmt.Println(maxNumber)
}