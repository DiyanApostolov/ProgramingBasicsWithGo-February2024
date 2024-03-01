package main

import (
	"fmt"
	"math"
)

//square, rectangle, circle или triangle

func main() {
	var figure string
	fmt.Scanln(&figure)

	if figure == "square" {
		var side float64
		fmt.Scanln(&side)

		var area float64 = side * side
		fmt.Printf("%.3f", area)

	} else if figure == "rectangle" {
		var sideA, sideB float64
		fmt.Scanln(&sideA)
		fmt.Scanln(&sideB)

		var area = sideA * sideB
		fmt.Printf("%.3f", area)
	} else if figure == "circle" {
		var radius float64
		fmt.Scanln(&radius)

		var area = math.Pi * math.Pow(radius, 2)
		fmt.Printf("%.3f", area)
	} else if figure == "triangle" {
		var sideA, hideA float64
		fmt.Scanln(&sideA)
		fmt.Scanln(&hideA)

		var area = sideA * hideA / 2
		fmt.Printf("%.3f", area)
	}
}
