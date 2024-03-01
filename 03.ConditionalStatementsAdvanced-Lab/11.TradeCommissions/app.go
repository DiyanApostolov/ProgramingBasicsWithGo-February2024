package main

import "fmt"

func main() {
	var city string
	var sales, commission float64

	fmt.Scanln(&city)
	fmt.Scanln(&sales)

	switch city {
	case "Sofia":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commission = 5
		} else if sales > 500 && sales <= 1000 {
			commission = 7
		} else if sales > 1000 && sales <= 10000 {
			commission = 8
		} else if sales > 10000 {
			commission = 12
		}
	case "Varna":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commission = 4.5
		} else if sales > 500 && sales <= 1000 {
			commission = 7.5
		} else if sales > 1000 && sales <= 10000 {
			commission = 10
		} else if sales > 10000 {
			commission = 13
		}
	case "Plovdiv":
		if sales < 0 {
			fmt.Println("error")
		} else if sales >= 0 && sales <= 500 {
			commission = 5.5
		} else if sales > 500 && sales <= 1000 {
			commission = 8
		} else if sales > 1000 && sales <= 10000 {
			commission = 12
		} else if sales > 10000 {
			commission = 14.5
		}
	default:
		fmt.Println("error")
	}

	if commission > 0 {
		output := sales * commission / 100

		fmt.Printf("%.2f", output)
	}

}
