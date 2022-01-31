package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, world")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello, this is the result of your request, your requested address:  %s\n", r.URL.Path)
	})

	http.ListenAndServe(":1290", nil)
}
