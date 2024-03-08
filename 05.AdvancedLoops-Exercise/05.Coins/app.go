package main

import "fmt"

func main() {
	var change float64
	fmt.Scanln(&change)
	var coinsCounter int

	var changeInStotinki int = int(change * 100)

	for changeInStotinki > 0 {
		if changeInStotinki >= 200 {
			changeInStotinki -= 200
			coinsCounter++
		} else if changeInStotinki >= 100{
			changeInStotinki -= 100
			coinsCounter++
		} else if changeInStotinki >= 50 {
			changeInStotinki -= 50
			coinsCounter++
		} else if changeInStotinki >= 20 {
			changeInStotinki -= 20
			coinsCounter++
		} else if changeInStotinki >= 10 {
			changeInStotinki -= 10
			coinsCounter++
		} else if changeInStotinki >= 5 {
			changeInStotinki -= 5
			coinsCounter++
		} else if changeInStotinki >= 2 {
			changeInStotinki -= 2
			coinsCounter++
		} else if changeInStotinki == 1 {
			changeInStotinki -= 1
			coinsCounter++
		} 
		
	}

	fmt.Println(coinsCounter)
}