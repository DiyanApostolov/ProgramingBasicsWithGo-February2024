package main

import "fmt"

func main() {
	var num, currentNum, p1, p2, p3, p4, p5 int
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&currentNum)

		if currentNum < 200 {
			p1++;
		} else if currentNum <= 399 {
			p2++;
		} else if currentNum <= 599 {
			p3++;
		} else if currentNum <= 799 {
			p4++;
		} else if currentNum >= 800 {
			p5++;
		}
	}

	percentP1 := float32(p1) / float32(num) * 100
	percentP2 := float32(p2) / float32(num) * 100
	percentP3 := float32(p3) / float32(num) * 100
	percentP4 := float32(p4) / float32(num) * 100
	percentP5 := float32(p5) / float32(num) * 100

	fmt.Printf("%.2f%%\n", percentP1)
	fmt.Printf("%.2f%%\n", percentP2)
	fmt.Printf("%.2f%%\n", percentP3)
	fmt.Printf("%.2f%%\n", percentP4)
	fmt.Printf("%.2f%%", percentP5)
}
