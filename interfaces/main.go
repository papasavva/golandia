package main

import "fmt"

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

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var shape1 Shape = Rectangle{
		width:  5,
		height: 10,
	}

	fmt.Println("Rectangle area:", shape1.Area())
	fmt.Println("Rectangle perimeter:", shape1.Perimeter())
}
