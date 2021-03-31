package main

import "fmt"

func main() {

	var luckyNumbers = []int8{7, 4, 9}

	for index, value := range luckyNumbers {
		fmt.Println("Index: ", index, " and Value: ", value)
	}
}
