package main

import "fmt"

func main() {
	var typeOfTicket string
	var rows, cols int
	var priceOfTicket float64

	fmt.Scan(&typeOfTicket, &rows, &cols)

	countOfTickets := rows * cols

	if typeOfTicket == "Premiere" {
		priceOfTicket = 12
	} else if typeOfTicket == "Normal" {
		priceOfTicket = 7.50
	} else if typeOfTicket == "Discount" {
		priceOfTicket = 5
	}

	finalIncome := float64(countOfTickets) * priceOfTicket

	fmt.Printf("%.2f leva", finalIncome)
}
