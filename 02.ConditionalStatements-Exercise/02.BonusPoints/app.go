package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	var bonusPoints float64

	if number > 1000 {
		bonusPoints = float64(number) * 0.1
	} else if number <= 100 {
		bonusPoints = 5
	} else if number > 100 {
		bonusPoints = float64(number) * 0.2
	}

	if number%2 == 0 {
		bonusPoints = bonusPoints + 1
	} else if number%10 == 5 {
		bonusPoints = bonusPoints + 2 //<- кратук запис   bonusPoints += 2
	}

	fmt.Println(bonusPoints)
	fmt.Println(float64(number) + bonusPoints)
}
