package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func hello() *int {
	i := 5
	return &i
}

func main() {
	x := 5
	double(&x)
	fmt.Println("Double of", x, "is ", x)

	d := hello()
	fmt.Println("Value of d", *d)
}
