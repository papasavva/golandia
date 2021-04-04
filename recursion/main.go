package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	const number = 6
	fmt.Println("Fibonacci of ", number, ":", fibonacci(number))
}
