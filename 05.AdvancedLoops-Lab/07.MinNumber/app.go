package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	var minNumber = math.MaxInt32

	for input != "Stop" {
		currentNum, _ := strconv.Atoi(input)
		
		if currentNum < minNumber{
			minNumber = currentNum
		}

		fmt.Scanln(&input)
	}

	fmt.Println(minNumber)
}