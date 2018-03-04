package main

import (
	"fmt"
)

//channel buffer

// unbuffer channel
// done <- true (sender)
// <-done (receiver)
func main() {

	// channel buffer,
	hello := make(chan string, 2)
	hello <- "Hello"
	hello <- "Golang"
	close(hello)

	for v := range hello {
		fmt.Println(v)
	}
}
