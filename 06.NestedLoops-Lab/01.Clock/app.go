package main

import "fmt"

func main() {

	for h := 0; h <= 23; h++ {
		for m := 0; m <= 59; m++ {
			fmt.Printf("%d:%d\n", h, m)
		}
	}
}
