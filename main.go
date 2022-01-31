package main

import "fmt"

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("Passed perfectly")
	} else if point > 5 {
		fmt.Println("Passed")
	} else if point == 4 {
		fmt.Println("Almost there")
	} else {
		fmt.Printf("You can do better! Your score is %d\n", point)
	}
}
