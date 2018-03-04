package main

import (
	"fmt"
)

func main() {

	p := &Person{
		ID:   1,
		Name: "Wuriyanto",
		Age:  26,
	}

	fmt.Println(p.GetName())
	p.ChangeName("Wuriyanto Musobar")
	fmt.Println(p.GetName())
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func (p *Person) GetID() int {
	return p.ID
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}
