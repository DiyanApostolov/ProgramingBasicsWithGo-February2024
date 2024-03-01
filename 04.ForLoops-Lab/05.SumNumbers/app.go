package main

import "fmt"

func main() {
	var n, currentNum, sum int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&currentNum)

		sum += currentNum
	}

	fmt.Println(sum)
}
