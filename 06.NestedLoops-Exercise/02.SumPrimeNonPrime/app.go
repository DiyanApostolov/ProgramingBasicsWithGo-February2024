package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sumPrimeNumbers, sumNonPrimeNumbers int
	var input string
	fmt.Scanln(&input)

	for input != "stop" {
		currentNumber, _ := strconv.Atoi(input)

		if currentNumber < 0 {
			fmt.Println("Number is negative.")
			fmt.Scanln(&input)
			continue
		}

		counter := 0

		for i := 1; i <= currentNumber; i++ {
			if currentNumber%i == 0 {
				counter++
			}
		}

		if counter == 2 {
			sumPrimeNumbers += currentNumber
		} else if counter > 2 {
			sumNonPrimeNumbers += currentNumber
		}

		fmt.Scanln(&input)
	}


	fmt.Printf("Sum of all prime numbers is: %d\n", sumPrimeNumbers)
	fmt.Printf("Sum of all non prime numbers is: %d", sumNonPrimeNumbers)
}
