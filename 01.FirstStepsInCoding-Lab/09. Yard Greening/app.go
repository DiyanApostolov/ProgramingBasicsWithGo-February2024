package main

import "fmt"

func main() {
	var yard float32

	fmt.Scanln(&yard)

	var price float32 = yard * 7.61
	var discount float32 = price * 0.18
	var finalPrice float32 = price - discount

	fmt.Printf("The final price is: %f lv.\n", finalPrice)
	fmt.Printf("The discount is: %f lv.", discount)
}
