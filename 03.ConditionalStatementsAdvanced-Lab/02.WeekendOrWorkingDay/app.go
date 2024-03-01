package main

import "fmt"

func main() {
	var day string
	fmt.Scanln(&day)

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Working day")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Error")
	}
}
