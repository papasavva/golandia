package main

import "fmt"

func main() {
	if 10%2 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	if dividend, divisor := 10, 2; dividend%divisor == 0 {
		fmt.Printf("Number %d is even", dividend)
	} else {
		fmt.Printf("Number %d is odd", dividend)

	}
}
