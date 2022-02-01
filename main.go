package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	var sayHelloTo = func(name string) {
		var data = fmt.Sprintf("Hello %s", name)
		messages <- data
	}

	go sayHelloTo("dude 1")
	go sayHelloTo("dude 2")
	go sayHelloTo("dude 3")

	var message1 = <-messages
	fmt.Println(message1)
	var message2 = <-messages
	fmt.Println(message2)
	var message3 = <-messages
	fmt.Println(message3)

}
