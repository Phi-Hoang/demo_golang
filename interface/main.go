package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	Width  float64
	Height float64
}

func (r Rect) Area() float64 {
	return 100
}
func (r Rect) Perimeter() float64 {
	return 100
}

type Circle struct {
	Diameter float64
}

func (r Circle) Area() float64 {
	return 200
}
func (r Circle) Perimeter() float64 {
	return 100
}
func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func PrintArea2(i interface{}) {
	// type assertion
	s, ok := i.(Shape)
	if ok {
		fmt.Println(s.Area())
	} else {
		fmt.Println("Not a shape")
	}
}

func main() {
	aRect := Rect{Height: 100, Width: 200}
	PrintArea(aRect)

	aCircle := Circle{Diameter: 300}
	PrintArea(aCircle)
	PrintArea2(aCircle)
}
