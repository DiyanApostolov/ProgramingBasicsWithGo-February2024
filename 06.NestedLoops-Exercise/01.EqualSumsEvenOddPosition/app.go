package main

import (
	"fmt"
	"strconv"
)

func main() {
	var startNumber, endNumber int
	fmt.Scanln(&startNumber)
	fmt.Scanln(&endNumber)

	for i := startNumber; i <= endNumber; i++ {
		var oddSum int = 0
		var evenSum int = 0

		// Обръщаме int число в string
		var currentNumber string = strconv.Itoa(i)

		for j := 0; j < 6; j++ {		
			currentDigit := currentNumber[j]	

			if j % 2 == 0 {
				evenSum += int(currentDigit) // кастваме към int стойността на текущата цифра
			} else {
				oddSum += int(currentDigit) // кастваме към int стойността на текущата цифра
			}
		}

		if oddSum == evenSum {
			fmt.Printf("%s ", currentNumber)
		}	
	}
}
