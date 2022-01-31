package main

import (
	"fmt"
	"strings"
)

func main() {
	var i = 0

	for i < 5 {
		fmt.Println(strings.Repeat("#", i))
		i++
	}

}
