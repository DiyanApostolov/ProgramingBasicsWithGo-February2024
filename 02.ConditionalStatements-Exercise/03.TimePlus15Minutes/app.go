package main

import "fmt"

func main() {
	var inputHour, inutMinutes int
	fmt.Scanln(&inputHour)
	fmt.Scanln(&inutMinutes)

	inutMinutes += 15

	var hour int = inputHour + inutMinutes / 60
	var minutes int = inutMinutes % 60

	if hour == 24 {
		hour = 0
	}

	fmt.Printf("%d:%.2d", hour, minutes)
}
