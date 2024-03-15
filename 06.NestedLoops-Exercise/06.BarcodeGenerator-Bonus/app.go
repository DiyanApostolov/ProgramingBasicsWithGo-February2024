package main

import "fmt"

func main() {
	var startNumber, endNumber int
	fmt.Scanln(&startNumber)
	fmt.Scanln(&endNumber)

	firstDigitStartNum := startNumber / 1000
	secondDigitStartNum := startNumber / 100 % 10
	thirdDigitStartNum := startNumber / 10 % 10
	forthDigitStartNum := startNumber % 10

	firstDigitEndNum := endNumber / 1000
	secondDigitEndNum := endNumber / 100 % 10
	thirdDigitEndNum := endNumber / 10 % 10
	forthDigitEndNum := endNumber % 10

	for x1 := firstDigitStartNum; x1 <= firstDigitEndNum; x1++ {
		for x2 := secondDigitStartNum; x2 <= secondDigitEndNum; x2++ {
			for x3 := thirdDigitStartNum; x3 <= thirdDigitEndNum; x3++ {
				for x4 := forthDigitStartNum; x4 <= forthDigitEndNum; x4++ {
					
					if x1 % 2 ==1 && x2 % 2 ==1 && x3 % 2 ==1 && x4 % 2 ==1{
						fmt.Printf("%d%d%d%d ", x1, x2, x3, x4)
					}

				}
			}
		}
	}
}
