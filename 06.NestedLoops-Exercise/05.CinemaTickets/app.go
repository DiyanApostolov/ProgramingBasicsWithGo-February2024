package main

import "fmt"

func main() {
	var nameOfMOvie, typeOfTicket string
	var seatsInTheater, ticketCounter, standardCounter, studentCounter, kidCounter, totalTickets int

	fmt.Scanln(&nameOfMOvie)

	for nameOfMOvie != "Finish" {
		fmt.Scanln(&seatsInTheater)
		fmt.Scanln(&typeOfTicket)
		ticketCounter = 0

		for typeOfTicket != "End" {
			switch typeOfTicket {
			case "standard":
				standardCounter++
				ticketCounter++
			case "student":
				studentCounter++
				ticketCounter++
			case "kid":
				kidCounter++
				ticketCounter++
			}

			if seatsInTheater == ticketCounter {
				break
			}
			
			fmt.Scanln(&typeOfTicket)
		}

        totalTickets += ticketCounter
		percentBusyTheater := float64(ticketCounter) / float64(seatsInTheater) * 100
		fmt.Printf("%s - %.2f%% full.\n", nameOfMOvie, percentBusyTheater)

		fmt.Scanln(&nameOfMOvie)
	}

	percentStandartTickets := float64(standardCounter) / float64(totalTickets) * 100
	percentStudentsTickets := float64(studentCounter) / float64(totalTickets) * 100
	percentKidsTickets := float64(kidCounter) / float64(totalTickets) * 100

	//output
	fmt.Printf("Total tickets: %d\n", totalTickets)
	fmt.Printf("%.2f%% student tickets.\n", percentStudentsTickets)
	fmt.Printf("%.2f%% standard tickets.\n", percentStandartTickets)
	fmt.Printf("%.2f%% kids tickets.", percentKidsTickets)

}
