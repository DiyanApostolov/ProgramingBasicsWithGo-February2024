package main

import "fmt"

func main() {
	for x := 1; x <= 10; x++ {
		for y := 1; y <= 10; y++ {
			result := x * y
			fmt.Printf("%d * %d = %d\n", x, y, result)
		}
	}
}
