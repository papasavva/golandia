package main

import (
	"errors"
	"fmt"
)

func main() {
	//var num float64 = 64
	sum, err := addPositiveNumbers(-1, 5)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum is:", sum)
	}

}

func addPositiveNumbers(number1 int, number2 int) (int, error) {
	if number1 < 0 || number2 < 0 {
		return 0, errors.New("Cannot add negative numbers")
	}

	return number1 + number2, nil
}
