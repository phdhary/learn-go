package main

import "fmt"

func main() {
	var point = 5

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}

}
