package main

import (
	"fmt"
)

func main() {
	var fruits = [...]string{"peach", "pear", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("fruit name: %s\n", fruit)
	}

}
