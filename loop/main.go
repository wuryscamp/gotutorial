package main

import (
	"fmt"
)

func main() {

	x := 1

	for {
		x++

		if x == 5 {
			continue
		}

		fmt.Printf("Hello go %d\n", x)

		if x == 10 {
			break
		}
	}

}
