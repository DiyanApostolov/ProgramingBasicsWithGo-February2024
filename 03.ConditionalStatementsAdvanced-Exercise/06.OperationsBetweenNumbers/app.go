package main

import "fmt"

func main() {
	var numOne, numTwo, result int
	var operator string

	fmt.Scan(&numOne, &numTwo, &operator)

	switch operator {
	case "+", "-", "*":
		if operator == "+" {
			result = numOne + numTwo
		} else if operator == "-" {
			result = numOne - numTwo
		} else if operator == "*" {
			result = numOne * numTwo
		}

		if result%2 == 0 {
			fmt.Printf("%d %s %d = %d - even", numOne, operator, numTwo, result)
		} else {
			fmt.Printf("%d %s %d = %d - odd", numOne, operator, numTwo, result)
		}
	case "/":
		if numTwo == 0 {
			fmt.Printf("Cannot divide %d by zero", numOne)
		} else {
			var resultAfterDivide float64 = float64(numOne) / float64(numTwo)
			
			fmt.Printf("%d / %d = %.2f", numOne, numTwo, resultAfterDivide)
		}
	case "%":
		if numTwo == 0 {
			fmt.Printf("Cannot divide %d by zero", numOne)
		} else {
			result = numOne % numTwo

			fmt.Printf("%d %% %d = %d", numOne, numTwo, result)
		}
	}
}
