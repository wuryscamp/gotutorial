package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		defer process(i)
	}
}

func process(id int) {
	fmt.Printf("Process %d\n", id)
}
