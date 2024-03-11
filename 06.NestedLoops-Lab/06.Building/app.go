package main

import "fmt"

func main() {
	var countOfFloors, countOfRooms int
	fmt.Scanln(&countOfFloors)
	fmt.Scanln(&countOfRooms)

	for floor := countOfFloors; floor >= 1; floor-- {
		for room := 0; room < countOfRooms; room++ {
			if floor == countOfFloors {
				// логика за последен етаж
				fmt.Printf("L%d%d ", floor, room)
			} else if floor%2 == 1 {
				// логика за нечетен етаж
				fmt.Printf("A%d%d ", floor, room)
			} else if floor%2 == 0 {
				// логика за четен етаж
				fmt.Printf("O%d%d ", floor, room)
			}
		}
		fmt.Println()
	}
}
