package main

import "fmt"

func main() {
	const grade uint = 11

	switch grade {
	case 20:
		fmt.Println("high")
	case 10, 11:
		fmt.Println("average")
	default:
		fmt.Println("other")
	}

	const temperature int = 20

	switch {
	case temperature > 30:
		fmt.Println("high")
	case temperature == 20:
		fmt.Println("medium")
	default:
		fmt.Println("low")
	}
}
