package main

import "fmt"

func main() {
	var point = 10000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s Passed perfectly", percent, "%")
	} else if point >= 70 {
		fmt.Printf("%.1f%s Good", percent, "%")
	} else {
		fmt.Printf("You can do better! Your score is %f\n", point)
	}
}
