package main

import "fmt"

func main() {
	// Constants
	const nylonPrice float32 = 1.50
	const paintPrice float32 = 14.50
	const distilerPrice float32 = 5
	const bagsPrice float32 = 0.40

	// Input
	var nylonInput, paintInput, distilerInput, hoursForWorking int

	fmt.Scanln(&nylonInput)
	fmt.Scanln(&paintInput)
	fmt.Scanln(&distilerInput)
	fmt.Scanln(&hoursForWorking)

	// Calculations

	var priceForNylon = (float32(nylonInput) + 2) * nylonPrice
	var priceForPaint = float32(paintInput) * 1.1 * paintPrice
	var priceForDistiler = float32(distilerInput) * distilerPrice
	var allPriceForMaterials = priceForNylon + priceForPaint + priceForDistiler + bagsPrice

	var priceForOneHour float32 = allPriceForMaterials * 0.3
	var priceForWork float32 = float32(hoursForWorking) * priceForOneHour

	var finalPrice = allPriceForMaterials + priceForWork

	// Ouput

	fmt.Println(finalPrice)
}
