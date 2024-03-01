package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "banana", "apple", "kiwi", "cherry", "lemon", "grapes":
		fmt.Println("fruit")
	case "tomato", "cucumber", "pepper", "carrot":
		fmt.Println("vegetable")
	default:
		fmt.Println("unknown")
	}
}
