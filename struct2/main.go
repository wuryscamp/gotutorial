package main

import (
	"fmt"
)

func main() {

	p := Person{
		ID:   1,
		Name: "Wuriyanto",
		Age:  26,
	}

	printPerson(p)
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func printPerson(p Person) {
	fmt.Println("ID = ", p.ID)
	fmt.Println("Name = ", p.Name)
	fmt.Println("Age = ", p.Age)
}
