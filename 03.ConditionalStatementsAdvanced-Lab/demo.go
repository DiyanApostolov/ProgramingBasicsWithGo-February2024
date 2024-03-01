package main

import "fmt"

func main() {

	var animal string
	fmt.Scanln(&animal)

	switch animal {
	case "dog":
		fmt.Println("Woof")
	case "cat":
		fmt.Println("Meow")
		fmt.Println("Meow")
		fmt.Println("Meow")
		fmt.Println("Meow")

	}
}
