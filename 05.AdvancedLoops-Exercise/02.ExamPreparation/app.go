package main

import (
	"fmt"
)

func main() {
	var maxPoorGrades, currentGrade, poorGradesCounter, problemsCounter, allGrades int
	var nameOfProblem, lastProblem string

	fmt.Scanln(&maxPoorGrades)
	fmt.Scanln(&nameOfProblem)

	for nameOfProblem != "Enough" {
		fmt.Scanln(&currentGrade)
		allGrades += currentGrade
		problemsCounter++

		if currentGrade <= 4 {
			poorGradesCounter++

			if poorGradesCounter == maxPoorGrades {
				fmt.Printf("You need a break, %d poor grades.", poorGradesCounter)
				break
			}
		}

		lastProblem = nameOfProblem

		fmt.Scanln(&nameOfProblem)
	}

	if nameOfProblem == "Enough" {
		averageScore := float32(allGrades) / float32(problemsCounter)

		fmt.Printf("Average score: %.2f\n", averageScore)
		fmt.Printf("Number of problems: %d\n", problemsCounter)
		fmt.Printf("Last problem: %s", lastProblem)
	}
}
