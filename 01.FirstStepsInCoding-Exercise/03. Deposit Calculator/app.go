package main

import "fmt"

func main() {
	var deposit float32
	var timeForDeposit int
	var percent float32

	fmt.Scanln(&deposit)
	fmt.Scanln(&timeForDeposit)
	fmt.Scanln(&percent)

	var profit float32 = deposit * percent / 100
	var profitForOneMonth float32 = profit / 12
	var finalProfit float32 = deposit + (float32(timeForDeposit) * profitForOneMonth)

	fmt.Println(finalProfit)
}
