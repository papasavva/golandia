package main

import "fmt"

func main() {
	//BASIC TYPES

	//numbers
	var level int = 15
	fmt.Println("level: ", level)

	var temperature int8 = -128
	fmt.Println("temperature overflow: ", temperature-1)

	//boolean
	var isLocked bool = false
	fmt.Println("front door is locked: ", isLocked)

	//String
	var favouriteRodent string = "hamster"
	fmt.Println("My favourite rodent is a ", favouriteRodent)

	//shorthand assignment statement in place of a var declaration with implicit type
	rodentSubFamily := "Cricetidae"
	fmt.Println("Rodent subfamilty: ", rodentSubFamily)

	var alwaysUtf8 string = "Go source files are always encoded in UTF-8\nand this is a proof.\nΕδώ γράφουμε Ελληνικα."
	fmt.Println(alwaysUtf8)

	//AGGREGATE (NON-REFERENCE) TYPES

	//arrays
	var favouriteNumbers [3]int = [3]int{6, 2, 9}
	fmt.Println(favouriteNumbers)

	var favouriteWords = [...]string{"excellent", "why"}
	fmt.Println(favouriteWords[1])

	//structs
	type customer struct {
		name string
		age  uint8
	}

	firstCustomer := customer{
		name: "Mickey Mouse",
		age:  12,
	}

	fmt.Println("First customer name: ", firstCustomer.name)
	fmt.Println("First customer age: ", firstCustomer.age)

	secondCustomer := customer{"Iron Man", 20}

	fmt.Println("Second customer: ", secondCustomer)
}
