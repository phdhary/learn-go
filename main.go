package main

import (
	"fmt"
)

func main() {
	var fruits = [...]string{"peach", "pear", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("Elements %d: %s\n", i, fruit)
	}

}
