package main

import (
	"fmt"
)

func main() {
	fmt.Println("fungsi main")

	x := 5
	y := 5

	z := add(x, y)
	fmt.Println(z)

	name := "Wuriyanto"
	result := hello(name)

	fmt.Println(result)

	xSwap := "Hello"
	ySwap := "Golang"

	resultX, resultY := swap(xSwap, ySwap)

	fmt.Println(resultX, resultY)
}

func add(x int, y int) int {
	return x + y
}

func hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func swap(x string, y string) (string, string) {
	return y, x
}
