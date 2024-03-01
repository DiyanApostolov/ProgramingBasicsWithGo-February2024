package main

import (
	"fmt"
	"math"
)

func main() {
	var num, currentNum, sum int
	var maxNum int32 = math.MinInt32

	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&currentNum)

		if int32(currentNum) > maxNum {
			maxNum = int32(currentNum)
		}

		sum += currentNum
	}

	sum -= int(maxNum)

	if sum == int(maxNum) {
		fmt.Println("Yes")
		fmt.Printf("Sum = %d", sum)
	} else {
		diff := math.Abs(float64(sum) - float64(maxNum))

		fmt.Println("No")
		fmt.Printf("Diff = %.0f", diff)
	}
}
