package main

import "fmt"

func main() {
	var destination string
	var budget, savedMoney, currentMoney float64

	fmt.Scanln(&destination)

	for destination != "End" {
		fmt.Scanln(&budget)

		for savedMoney < budget {
			fmt.Scanln(&currentMoney)
			savedMoney += currentMoney
		}

		fmt.Printf("Going to %s!\n", destination)
		savedMoney = 0

		fmt.Scanln(&destination)
	}
}
