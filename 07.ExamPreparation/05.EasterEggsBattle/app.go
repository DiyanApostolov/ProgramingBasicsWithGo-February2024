package main

import (
	"fmt"
)

func main() {
	var firstPlayerEggs, secondPlayerEggs int
	var input string
	fmt.Scanln(&firstPlayerEggs)
	fmt.Scanln(&secondPlayerEggs)
	fmt.Scanln(&input)

	for input != "End" {
		if input == "one" {
			secondPlayerEggs--
		} else if input == "two" {
			firstPlayerEggs--
		}

		if firstPlayerEggs == 0 {
			fmt.Printf("Player one is out of eggs. Player two has %d eggs left.", secondPlayerEggs)
			break
		}
		if secondPlayerEggs == 0 {
			fmt.Printf("Player two is out of eggs. Player one has %d eggs left.", firstPlayerEggs)
			break
		}

		fmt.Scanln(&input)
	}

	if input == "End" {
		fmt.Printf("Player one has %d eggs left.\n", firstPlayerEggs)
		fmt.Printf("Player two has %d eggs left.", secondPlayerEggs)
	}
}
