package main

import (
	"fmt"
)

func main() {
	//"Roses", "Dahlias", "Tulips", "Narcissus", "Gladiolus"

	const priceForOneRose float64 = 5
	const priceForOneDahlia float64 = 3.80
	const priceForOneTulip float64 = 2.80
	const priceForOneNarcissus float64 = 3
	const priceForOneGladiolus float64 = 2.5

	var typeOfFlower string
	var countOfFlowers, budget int
	var finalPrice float64

	fmt.Scan(&typeOfFlower, &countOfFlowers, &budget)

	if typeOfFlower == "Roses" {
		finalPrice = float64(countOfFlowers) * priceForOneRose

		if countOfFlowers > 80 {
			finalPrice *= 0.9
		}
	} else if typeOfFlower == "Dahlias" {
		finalPrice = float64(countOfFlowers) * priceForOneDahlia

		if countOfFlowers > 90 {
			finalPrice *= 0.85
		}
	} else if typeOfFlower == "Tulips" {
		finalPrice = float64(countOfFlowers) * priceForOneTulip

		if countOfFlowers > 80 {
			finalPrice *= 0.85
		}
	} else if typeOfFlower == "Narcissus" {
		finalPrice = float64(countOfFlowers) * priceForOneNarcissus

		if countOfFlowers < 120 {
			finalPrice *= 1.15
		}
	} else if typeOfFlower == "Gladiolus" {
		finalPrice = float64(countOfFlowers) * priceForOneGladiolus

		if countOfFlowers < 80 {
			finalPrice *= 1.2
		}
	}

	if float64(budget) >= finalPrice {
		diff := float64(budget) - finalPrice
		fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", countOfFlowers, typeOfFlower, diff)
	} else {
		neededMoney := finalPrice - float64(budget)
		fmt.Printf("Not enough money, you need %.2f leva more.", neededMoney)
	}

}
