package main

import "fmt"

func main() {
	var luckyNumbers = []int{1, 2, 3, 6}
	fmt.Println("Lucky numbers sum:", sum(luckyNumbers...))
}

func sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
