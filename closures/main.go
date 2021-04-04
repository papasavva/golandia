package main

import "fmt"

func counter() func() int {
	current := 0
	return func() int {
		current += 1
		return current
	}
}

func main() {
	firstCounter := counter()
	fmt.Println("Current index:", firstCounter())
	fmt.Println("Current index:", firstCounter())
	fmt.Println("Current index:", firstCounter())

	//we can confirm that a new variable has unique state
	newCounter := counter()
	fmt.Println("Current index:", newCounter())
	fmt.Println("Current index:", newCounter())
	fmt.Println("Current index:", newCounter())
}
