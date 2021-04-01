package main

import "fmt"

func main() {

	var luckyNumbers = []int8{7, 4, 9}

	for index, value := range luckyNumbers {
		fmt.Println("Index:", index, " and Value:", value)
	}

	var colours = make(map[string]string)
	colours["red"] = "#FF0000"
	colours["blue"] = "#0000FF"
	colours["yellow"] = "#FFFF00"

	for index, value := range colours {
		fmt.Println("Colour:", index, " and Hex Code:", value)
	}
}
