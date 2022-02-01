package main

import (
	"fmt"
)

func main() {
	var number = 4
	fmt.Println("before: ", number)

	change(&number, 50)
	fmt.Println("after: ", number)

}

func change(original *int, value int) {
	*original = value
}
