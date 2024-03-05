package main

import "fmt"

func main() {
	var username, password, input string

	fmt.Scanln(&username)
	fmt.Scanln(&password)

	for true {
		fmt.Scanln(&input)

		if input == password {
			fmt.Printf("Welcome %s!", username)
			break
		}
	}
}
