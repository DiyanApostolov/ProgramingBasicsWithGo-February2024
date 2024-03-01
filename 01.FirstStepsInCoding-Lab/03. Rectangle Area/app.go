package main

import "fmt"

func main() {
	var sideA int
	var sideB int

	fmt.Scanln(&sideA)
	fmt.Scanln(&sideB)

	var result int = sideA * sideB

	fmt.Println(result)
}
