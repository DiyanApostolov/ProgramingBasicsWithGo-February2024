package main

import "fmt"

func main() {
	var age, priceForToy, toyCounter, lilyMoney int
	var priceForWashingMachine float64
	var birthdayMoney = 10

	fmt.Scanln(&age)
	fmt.Scanln(&priceForWashingMachine)
	fmt.Scanln(&priceForToy)

	for i := 1; i <= age; i++ {

		if i%2 != 0 { //нечетно число
			toyCounter++
		} else { //четно число
			lilyMoney += birthdayMoney
			lilyMoney--

			birthdayMoney += 10
		}
	}

	lilyMoney += toyCounter * priceForToy

	if float64(lilyMoney) >= priceForWashingMachine {
		diff := float64(lilyMoney) - priceForWashingMachine

		fmt.Printf("Yes! %.2f", diff)
	} else {
		moneyNeeded := priceForWashingMachine - float64(lilyMoney)

		fmt.Printf("No! %.2f", moneyNeeded)
	}

}
