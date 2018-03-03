package main

import (
	"fmt"
)

func main() {

	add := func(x, y int) int {
		return x + y
	}

	hello := func(name string) string {
		return fmt.Sprintf("Hello %s", name)
	}

	fmt.Println(add(5, 5))
	fmt.Println(hello("Wuriyanto"))
}
