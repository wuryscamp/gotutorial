package main

import (
	"fmt"
)

//channel
func main() {

	done := make(chan bool)

	go helloGo(done)

	<-done
	fmt.Println("Ini Fungsi main")
}

func helloGo(done chan bool) {
	fmt.Println("Hello Go")
	done <- true
}
