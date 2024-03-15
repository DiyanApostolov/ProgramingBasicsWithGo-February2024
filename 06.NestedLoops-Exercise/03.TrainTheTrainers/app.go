package main

import "fmt"

func main() {
	var countOfGrades, presentationCounter int
	var nameOfPresentation string
	fmt.Scanln(&countOfGrades)
	fmt.Scanln(&nameOfPresentation)

	var currentGrade, sumAllAverageGrades float64

	for nameOfPresentation != "Finish" {
		presentationCounter++
		sumOfPresentationGrades := 0.0

		for i := 0; i < countOfGrades; i++ {
			fmt.Scanln(&currentGrade)
			sumOfPresentationGrades += currentGrade
		}

		averageGrade := sumOfPresentationGrades / float64(countOfGrades)
		sumAllAverageGrades += averageGrade
		fmt.Printf("%s - %.2f.\n", nameOfPresentation, averageGrade)

		fmt.Scanln(&nameOfPresentation)
	}

	finalAssessment := sumAllAverageGrades / float64(presentationCounter)

	fmt.Printf("Student's final assessment is %.2f.", finalAssessment)
}
