package main

import (
	"fmt"
)

func main() {
	var names [4]string

	names[0] = "john"
	names[1] = "deto"
	names[2] = "berto"
	names[3] = "bernado"

	fmt.Println(names, names[3])
}
