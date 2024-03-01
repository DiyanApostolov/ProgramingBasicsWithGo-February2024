package main

import "fmt"

func main() {
	var dergees int
	var timeOfDay, outfit, shoes string

	fmt.Scan(&dergees, &timeOfDay)

	if dergees >= 10 && dergees <= 18 {
		switch timeOfDay {
		case "Morning":
			outfit = "Sweatshirt"
			shoes = "Sneakers"
		case "Afternoon", "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
		}
	} else if dergees > 18 && dergees <= 24 {
		switch timeOfDay {
		case "Morning", "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
		case "Afternoon":
			outfit = "T-Shirt"
			shoes = "Sandals"

		}
	} else if dergees > 24 {
		switch timeOfDay {
		case "Morning":
			outfit = "T-Shirt"
			shoes = "Sandals"
		case "Afternoon":
			outfit = "Swim Suit"
			shoes = "Barefoot"
		case "Evening":
			outfit = "Shirt"
			shoes = "Moccasins"
		}
	}

	fmt.Printf("It's %d degrees, get your %s and %s.", dergees, outfit, shoes)

}
