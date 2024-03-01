package main

import "fmt"

func main() {
	var inchesInput float64

	fmt.Scanln(&inchesInput)

	var centimeters float64 = inchesInput * 2.54

	fmt.Println(centimeters)
}
