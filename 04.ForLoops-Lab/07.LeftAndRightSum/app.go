package main

import (
	"fmt"
	"math"
)

func main() {
	var n, leftSum, rigthSum, currentNum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)
		leftSum += currentNum
	}

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)
		rigthSum += currentNum
	}

	if leftSum == rigthSum {
		fmt.Printf("Yes, sum = %d", leftSum)
	} else {
		difference := math.Abs(float64(leftSum) - float64(rigthSum))

		fmt.Printf("No, diff = %.0f", difference)
	}
}
