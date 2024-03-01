package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 0; i <= num; i+=2 {
		fmt.Println(math.Pow(2, float64(i)))
	}
}
