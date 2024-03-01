package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var age int
	var town string

	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)
	fmt.Scanln(&age)
	fmt.Scanln(&town)

	fmt.Printf("You are %s %s, a %d-years old person from %s.", firstName, lastName, age, town)
}
