package main

import "fmt"

func main() {
	// season "summer" или "winter"
	// destinations "Bulgaria", "Balkans" и "Europe"
	// място за отсядане "Camp" и "Hotel"

	var budget, spendMoney float64
	var season, destination, placeForRest string

	fmt.Scan(&budget, &season)

	if budget <= 100 {
		destination = "Bulgaria"

		switch season {
		case "summer":
			spendMoney = budget * 0.3
			placeForRest = "Camp"
		case "winter":
			spendMoney = budget * 0.7
			placeForRest = "Hotel"
		}
	} else if budget <= 1000 {
		destination = "Balkans"

		switch season {
		case "summer":
			spendMoney = budget * 0.4
			placeForRest = "Camp"
		case "winter":
			spendMoney = budget * 0.8
			placeForRest = "Hotel"
		}
	} else if budget > 1000 {
		destination = "Europe"
		placeForRest = "Hotel"
		spendMoney = budget * 0.9
	}

	fmt.Printf("Somewhere in %s\n", destination)
	fmt.Printf("%s - %.2f", placeForRest, spendMoney)

}
