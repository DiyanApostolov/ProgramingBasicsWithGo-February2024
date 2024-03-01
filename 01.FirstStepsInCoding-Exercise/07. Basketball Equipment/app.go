package main

import "fmt"

func main() {
	// Input
	var pricePerYear int
	fmt.Scanln(&pricePerYear)

	// Calculations
	var priceForSneakers float32 = float32(pricePerYear) - (float32(pricePerYear) * 0.4) // pricePerYear * 0.6
	var priceForClothes float32 = priceForSneakers * 0.8
	var priceForBall float32 = priceForClothes * 0.25
	var priceForAccessories float32 = priceForBall * 0.2

	var finalPrice = priceForSneakers + priceForClothes + priceForBall + priceForAccessories + float32(pricePerYear)

	//Output

	fmt.Println(finalPrice)
}
