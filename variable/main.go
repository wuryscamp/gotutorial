package main

import (
	"fmt"
)

func main() {

	// static type declaration
	var x int
	x = 10

	var y float64
	y = 5.5

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("x memiliki tipe data %T\n", x)
	fmt.Printf("y memiliki tipe data %T\n", y)

	// dynamic type declaration
	z := "wuriyanto"
	i := 50
	fmt.Println(z)
	fmt.Println(i)
	fmt.Printf("z memiliki tipe data %T\n", z)
	fmt.Printf("i memiliki tipe data %T\n", i)

	t := 5
	p := 10

	fmt.Println(t + p)
}
