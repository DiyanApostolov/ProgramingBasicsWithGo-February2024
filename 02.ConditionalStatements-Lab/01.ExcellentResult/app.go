package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	if input >= 5 {
		fmt.Println("Excellent!")
	}
}
