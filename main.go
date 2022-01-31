package main

import (
	"fmt"
)

func main() {
	var firstName string = "Berto"

	lastName := "Chagreta"

	_ = "this is variable underscore that will be ignored"
	_ = "this is variable underscore that will be ignored too"

	fruit, _ := "apple", "grape"

	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Print(fruit)
}
