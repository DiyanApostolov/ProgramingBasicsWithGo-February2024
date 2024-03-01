package main

import "fmt"

func main() {
	var countOfTabs, salary int
	var tabName string

	fmt.Scanln(&countOfTabs)
	fmt.Scanln(&salary)

	for i := 0; i < countOfTabs; i++ {
		fmt.Scanln(&tabName)

		switch tabName {
		case "Facebook":
			salary -= 150
		case "Instagram":
			salary -= 100
		case "Reddit":
			salary -= 50
		}

		if salary <= 0 {
			break
		}
	}

	if salary > 0 {
		fmt.Println(salary)
	} else {
		fmt.Println("You have lost your salary.")
	}
}
