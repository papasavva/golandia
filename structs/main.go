package main

import "fmt"

type Customer struct {
	name string
	age  int
}

func main() {
	customer1 := Customer{
		name: "Bob",
		age:  20,
	}

	fmt.Println("Customer:", customer1)
}
