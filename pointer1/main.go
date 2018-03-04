package main

import (
	"fmt"
)

func main() {

	var hello string = "Hello"
	var strPtr *string

	strPtr = &hello

	fmt.Println(&hello)
	fmt.Println(strPtr)

	fmt.Println(hello)
	changePtr(&hello)
	fmt.Println(hello)
}

// pass by value
func change(v string) {
	v = v + " Golang"
}

// pass by reference
func changePtr(v *string) {
	*v = *v + " Golang"
}
