package main

import (
	"fmt"
	"strings"
)

func main() {
	var i = 0

	for {
		fmt.Println(strings.Repeat("#", i))
		i++
		if i == 10 {
			break
		}
	}

}
