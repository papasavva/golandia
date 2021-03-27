package main

import "fmt"

func main() {
	//BASIC TYPES

	//numbers
	var level int = 15
	fmt.Println("level: ", level)

	var temperature int8 = -128
	fmt.Println("temperature overflow: ", temperature-1)
}
