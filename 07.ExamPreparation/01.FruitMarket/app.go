package main

import "fmt"

func main() {
	var priceOfStrawberry, amounOfBanana, amountOfOrage, amountOfRasberry, amoutOfStrawberry float64

	//Input
	fmt.Scanln(&priceOfStrawberry)
	fmt.Scanln(&amounOfBanana)
	fmt.Scanln(&amountOfOrage)
	fmt.Scanln(&amountOfRasberry)
	fmt.Scanln(&amoutOfStrawberry)

	// Calculations
	priceOfRasberry := priceOfStrawberry / 2
	priceOfOrange := priceOfRasberry * 0.6
	priceOfBanana := priceOfRasberry * 0.2

	finalPrice := (priceOfStrawberry * amoutOfStrawberry) + (priceOfBanana * amounOfBanana) + (priceOfOrange * amountOfOrage) + (priceOfRasberry * amountOfRasberry)

	//Output
	fmt.Printf("%.2f", finalPrice)
}
