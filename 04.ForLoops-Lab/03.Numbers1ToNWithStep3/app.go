package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 1; i <= num; i+=3 {
		fmt.Println(i)
	}
}
