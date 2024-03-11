package main

import "fmt"

func main() {
	var start, end, magicNumber, counter int
	fmt.Scan(&start, &end, &magicNumber)

	var isFound bool = false


	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			counter++
			sum := i + j

			if sum == magicNumber {
				fmt.Printf("Combination N:%d (%d + %d = %d)\n", counter, i, j, magicNumber)
				isFound = true
				break
			}
		}

		if isFound {
			break
		}
	}

	if !isFound {
		fmt.Printf("%d combinations - neither equals %d", counter, magicNumber)
	}
}
