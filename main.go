package main

import "fmt"

func main() {
	// new keyword to create var with type's default pointer value
	name := new(string)
	fmt.Println(name)
	// to access it's value use asterisk *
	fmt.Println(*name)
}
