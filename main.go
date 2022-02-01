package main

import (
	"fmt"
)

func main() {
	var fruits = [...]string{"peach", "pear", "banana", "melon"}

	fmt.Println("fruits count: ", len(fruits))
	fmt.Println("fruits: ", fruits)
}
