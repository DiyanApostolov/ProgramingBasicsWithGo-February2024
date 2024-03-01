package main

import "fmt"

func main() {
	var arrivalHour, arrivalMinutes, examHour, examMinutes int

	fmt.Scan(&examHour, &examMinutes, &arrivalHour, &arrivalMinutes)

	examTimeInMInutes := examHour*60 + examMinutes
	arrivalTimeInMInutes := arrivalHour*60 + arrivalMinutes

	if arrivalTimeInMInutes > examTimeInMInutes {

		fmt.Println("Late")

		diff := arrivalTimeInMInutes - examTimeInMInutes

		if diff < 60 {
			fmt.Printf("%d minutes after the start", diff)
		} else {
			lateHour := diff / 60
			lateMinutes := diff % 60

			fmt.Printf("%d:%.2d hours after the start", lateHour, lateMinutes)
		}
	} else if examTimeInMInutes-arrivalTimeInMInutes <= 30 || examTimeInMInutes == arrivalTimeInMInutes {
		fmt.Println("On time")

		if examTimeInMInutes != arrivalTimeInMInutes {
			diff := examTimeInMInutes - arrivalTimeInMInutes

			fmt.Printf("%d minutes before the start", diff)
		}
	} else if examTimeInMInutes-arrivalTimeInMInutes > 30 {
		fmt.Println("Early")

		diff := examTimeInMInutes - arrivalTimeInMInutes

		if diff < 60 {
			fmt.Printf("%d minutes before the start", diff)
		} else {
			earlyHour := diff / 60
			earlyMinutes := diff % 60

			fmt.Printf("%d:%.2d hours before the start", earlyHour, earlyMinutes)
		}
	}
}
