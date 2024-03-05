package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// initialzing scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	var banckAccount float64
	var input string
	
	//read input
	scanner.Scan()         // scan one row frow console
	input = scanner.Text() // send text from console to var input

	for input != "NoMoreMoney" {
		currentAmount, _ := strconv.ParseFloat(input, 64)

		if currentAmount < 0 {
			fmt.Println("Invalid operation!")
			break
		}

		banckAccount += currentAmount

		fmt.Printf("Increase: %.2f\n", currentAmount)

		scanner.Scan()         // scan one row frow console
		input = scanner.Text() // send text from console to var input
	}

	fmt.Printf("Total: %.2f", banckAccount)
}
