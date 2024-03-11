package main

import "fmt"

func main() {
	var num, validCombinations int
	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		for j := 0; j <= num; j++ {
			for k := 0; k <= num; k++ {
				sum := i + j + k

				if sum == num {
					validCombinations++
				}
			}
		}
	}

	fmt.Println(validCombinations)
}
