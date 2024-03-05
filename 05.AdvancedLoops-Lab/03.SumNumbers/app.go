package main

import "fmt"

func main() {
	var number, input, sum int
	fmt.Scanln(&number)

	for sum < number {
		fmt.Scanln(&input)
		sum += input
	}

	fmt.Println(sum)
}
