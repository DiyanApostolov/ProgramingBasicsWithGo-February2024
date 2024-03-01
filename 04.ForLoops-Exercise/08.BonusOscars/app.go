package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var name string
	var pointsFromAcademy float64
	var countOfPointers int

	scanner.Scan()
	name = scanner.Text()
	fmt.Scanln(&pointsFromAcademy)
	fmt.Scanln(&countOfPointers)

	for i := 0; i < countOfPointers; i++ {
		var nameOfPointer string
		var currentPoints float64
		scanner.Scan()
		nameOfPointer = scanner.Text()
		fmt.Scanln(&currentPoints)

		points := float64(len(nameOfPointer)) * currentPoints / 2
		pointsFromAcademy += points

		if pointsFromAcademy >= 1250.5 {
			break
		}
	}

	if pointsFromAcademy >= 1250.5 {
		fmt.Printf("Congratulations, %s got a nominee for leading role with %.1f!", name, pointsFromAcademy)
	} else {
		neededPoints := 1250.5 - pointsFromAcademy
		fmt.Printf("Sorry, %s you need %.1f more!", name, neededPoints)
	}

}
