package main

import "fmt"

func main() {
	var country, sovenir string
	var countOfSovenis int
	var price float64
	fmt.Scan(&country, &sovenir, &countOfSovenis)

	if country == "Argentina" {
		switch sovenir {
		case "flags":
			price = 3.25
		case "caps":
			price = 7.20
		case "posters":
			price = 5.10
		case "stickers":
			price = 1.25
		default:
			fmt.Println("Invalid stock!")
		}
	} else if country == "Brazil" {
		switch sovenir {
		case "flags":
			price = 4.20
		case "caps":
			price = 8.50
		case "posters":
			price = 5.35
		case "stickers":
			price = 1.20
		default:
			fmt.Println("Invalid stock!")
		}
	} else if country == "Croatia" {
		switch sovenir {
		case "flags":
			price = 2.75
		case "caps":
			price = 6.90
		case "posters":
			price = 4.95
		case "stickers":
			price = 1.10
		default:
			fmt.Println("Invalid stock!")
		}
	} else if country == "Denmark" {
		switch sovenir {
		case "flags":
			price = 3.10
		case "caps":
			price = 6.50
		case "posters":
			price = 4.80
		case "stickers":
			price = 0.90
		default:
			fmt.Println("Invalid stock!")
		}
	} else {
		fmt.Println("Invalid country!")
	}

	if price > 0 {
		finalPrice := float64(countOfSovenis) * price

		fmt.Printf("Pepi bought %d %s of %s for %.2f lv.", countOfSovenis, sovenir, country, finalPrice)
	}

}
