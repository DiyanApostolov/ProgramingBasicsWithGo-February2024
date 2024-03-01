package main

import "fmt"

func main() {
	var fisrtSportist, secondSportist, thirdSportist int
	fmt.Scanln(&fisrtSportist)
	fmt.Scanln(&secondSportist)
	fmt.Scanln(&thirdSportist)

	var sum int = fisrtSportist + secondSportist + thirdSportist

	var minutes int = sum / 60
	var seconds int = sum % 60

	if seconds < 10 {
		fmt.Printf("%d:0%d", minutes, seconds)
	} else {
		fmt.Printf("%d:%d", minutes, seconds)
	}
}
