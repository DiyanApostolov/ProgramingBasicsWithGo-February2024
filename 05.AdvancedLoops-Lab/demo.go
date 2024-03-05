package main

import "fmt"

func main() {

	var num = 0

	for num <= 10 {
		num++

		if num%2 == 1 {
			continue
		}

		fmt.Println(num)
	}
}
