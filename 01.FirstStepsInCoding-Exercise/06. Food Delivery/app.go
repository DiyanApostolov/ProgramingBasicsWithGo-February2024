package main

import "fmt"

func main() {
	// Input
	var countOfChikenMenus int
	var countOfFishMenus int
	var countOfVegetarianMenus int

	fmt.Scanln(&countOfChikenMenus)
	fmt.Scanln(&countOfFishMenus)
	fmt.Scanln(&countOfVegetarianMenus)

	// Calculations
	var priceForChikenMenus float32 = float32(countOfChikenMenus) * 10.35
	var priceForFishMenus float32 = float32(countOfFishMenus) * 12.40
	var priceForVegetarianMenus float32 = float32(countOfVegetarianMenus) * 8.15

	var allPurchase = priceForChikenMenus + priceForFishMenus + priceForVegetarianMenus
	var dessert = allPurchase * 0.2
	var finalPrice = allPurchase + dessert + 2.50

	// Output

	fmt.Println(finalPrice)

}
