package main

import "fmt"

func main() {
	var end int
	fmt.Scanln(&end)

	currentNum := 1

	for currentNum <= end {
		fmt.Println(currentNum)

		currentNum = currentNum * 2 + 1
	}
}