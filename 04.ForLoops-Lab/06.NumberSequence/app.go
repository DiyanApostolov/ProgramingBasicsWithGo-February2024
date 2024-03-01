package main

import (
	"fmt"
	"math"
)

func main() {
	var n, currentNum int
	fmt.Scanln(&n)

	var minNumber int32 = math.MaxInt32
	var maxNumber int32 = math.MinInt32

	for i := 0; i < n; i++ {
		fmt.Scanln(&currentNum)
			
		if int32(currentNum) < minNumber{
			minNumber = int32(currentNum)
		}

		if int32(currentNum) > maxNumber {
			maxNumber = int32(currentNum)
		}
	}

	fmt.Printf("Max number: %d\n", maxNumber)
	fmt.Printf("Min number: %d", minNumber)
}
