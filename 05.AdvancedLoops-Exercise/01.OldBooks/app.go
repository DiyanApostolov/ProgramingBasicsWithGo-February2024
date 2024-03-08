package main

import "fmt"

func main() {
	var favoriteBook, nextBook string
	var bookCounter int
	var isFound bool = false

	fmt.Scanln(&favoriteBook)
	fmt.Scanln(&nextBook)

	for nextBook != "NoMoreBooks" {
		if favoriteBook == nextBook {
			fmt.Printf("You checked %d books and found it.", bookCounter)
			isFound = true
			break
		}

		bookCounter++

		fmt.Scanln(&nextBook)
	}

	if !isFound {
		fmt.Println("The book you search is not here!")
		fmt.Printf("You checked %d books.", bookCounter)
	}
	
}
