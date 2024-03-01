package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := num; i>=1; i--{
		fmt.Println(i)
	}
}