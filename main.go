package main

import "fmt"

func main() {
	var point = 5

	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}

}
