package main

import "fmt"

func main() {
	var decimalNumber = 2.62
	// %f to format decimal to string
	// %3.f to format 3 digit decimal to string
	fmt.Printf("decimal num: %f\n", decimalNumber)
	fmt.Printf("decimal num: %.3f\n", decimalNumber)
}
