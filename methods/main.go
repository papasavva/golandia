package main

import "fmt"

type Customer struct {
	name string
	age  int
}

//method with a reciver of customer type
func (c Customer) print() {
	fmt.Println("Customer name", c.name)
	fmt.Println("Customer age", c.age)
}

func main() {
	firstCustomer := Customer{"Patricio", 20}
	firstCustomer.print()
}
