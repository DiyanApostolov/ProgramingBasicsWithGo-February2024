package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for x1 := 1; x1 <= 9; x1++ {
		for x2 := 1; x2 <= 9; x2++ {
			for x3 := 1; x3 <= 9; x3++ {
				for x4 := 1; x4 <= 9; x4++ {

					if num % x1 == 0 && num % x2 == 0 && num % x3 == 0 && num % x4 == 0{
						fmt.Printf("%d%d%d%d ", x1, x2, x3, x4)
					}
					
				}
			}
		}
	}
}