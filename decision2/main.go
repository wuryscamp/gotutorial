package main

import (
	"fmt"
)

func main() {

	x := 100

	if x < 10 {
		fmt.Println("Hello go 1")
	} else if x >= 100 {
		fmt.Println("Hello go 2")
	} else {
		fmt.Println("Hello go 3")
	}
}
