package main

import "fmt"

func double(number *int) {
	*number *= 2
}
func main() {
	x := 5
	double(&x)
	fmt.Println("Double of", x, "is ", x)
}
