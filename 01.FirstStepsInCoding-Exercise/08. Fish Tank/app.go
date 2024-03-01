package main

import "fmt"

func main() {
	// Input
	var sideA, sideB, sideC int
	var percent float32

	fmt.Scanln(&sideA)
	fmt.Scanln(&sideB)
	fmt.Scanln(&sideC)
	fmt.Scanln(&percent)

	// Calculations
	var aquariumCumSantimetres int = sideA * sideB * sideC
	var aquariumLitres float32 = float32(aquariumCumSantimetres) * 0.001
	var neededLitres float32 = aquariumLitres * (1 - (percent / 100))

	// Output
	fmt.Println(neededLitres)
}
