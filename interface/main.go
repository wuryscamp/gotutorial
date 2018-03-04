package main

import (
	"fmt"
	"math"
)

func main() {

	circle := Circle{X: 1.0, Y: 1.0, Radius: 10.5}
	rectangle := Rectangle{Width: 10, Height: 5}

	fmt.Println("Area dari Circle = ", getArea(circle))
	fmt.Println("Area dari Rectangle = ", getArea(rectangle))
}

type Shape interface {
	Area() float64
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func getArea(s Shape) float64 {
	return s.Area()
}
