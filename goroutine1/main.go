package main

import (
	"fmt"
	"time"
)

func main() {

	go helloGo()

	time.Sleep(1 * time.Second)
	fmt.Println("Ini Fungsi main")
}

func helloGo() {
	fmt.Println("Hello Go")
}
