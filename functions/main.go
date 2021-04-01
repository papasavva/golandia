package main

import "fmt"

func main() {
	fmt.Println("Circle Area:", circleArea(3))
	fmt.Println("Number is odd:", numberIsOdd(3))

	sum, average := sumAndAverage(9, 3)
	fmt.Println("sum:", sum, " average: ", average)
}

func circleArea(radius float32) float32 {
	const PI float32 = 3.14

	return PI * radius * radius
}

func numberIsOdd(number int) bool {
	return number%2 != 0
}

func sumAndAverage(a int, b int) (int, int) {
	sum := a + b
	average := sum / 2

	return sum, average
}
