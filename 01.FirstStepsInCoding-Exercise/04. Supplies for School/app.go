package main

import "fmt"

func main() {
	const priceForPens float32 = 5.80
	const priceForMarkers float32 = 7.20
	const priceForPreparat float32 = 1.20

	var countOfPensPackage, countOfMarkersPackage, littresOfPreparat, percentOfDiscount int

	fmt.Scanln(&countOfPensPackage)
	fmt.Scanln(&countOfMarkersPackage)
	fmt.Scanln(&littresOfPreparat)
	fmt.Scanln(&percentOfDiscount)

	var priceForAllPens float32 = float32(countOfPensPackage) * priceForPens
	var priceForAllMarkers float32 = float32(countOfMarkersPackage) * priceForMarkers
	var priceForAllPreparat float32 = float32(littresOfPreparat) * priceForPreparat

	var price = priceForAllPens + priceForAllMarkers + priceForAllPreparat
	var discount = price * float32(percentOfDiscount) / 100

	var finalPrice = price - discount

	fmt.Println(finalPrice)
}
