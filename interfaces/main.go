package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

//Rectangle
func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

//Circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var shape1 Shape = Rectangle{
		width:  5,
		height: 10,
	}

	var shape2 Shape = Circle{5}

	fmt.Println("Rectangle area:", shape1.Area())
	fmt.Println("Rectangle perimeter:", shape1.Perimeter())

	fmt.Println("Circle area:", shape2.Area())
	fmt.Println("Circle perimeter:", shape2.Perimeter())
}
