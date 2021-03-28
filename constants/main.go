package main

import "fmt"

func main() {
	const TAX_PERCENTAGE = 10
	fmt.Println(TAX_PERCENTAGE)

	//untyped constant
	const AUTHORITY = "hacienda"
	const AUTHORITY_TYPED string = "hacienda (typed)"
	fmt.Println(AUTHORITY)

	var s string
	s = AUTHORITY       // This works just fine
	s = AUTHORITY_TYPED // As does this
	fmt.Println(s)

	type MyString string
	var mys MyString
	// This works just fine
	mys = AUTHORITY
	// cannot use AUTHORITY_TYPED (type string) as type MyString in assignment
	//mys = AUTHORITY_TYPED

	fmt.Println(mys)

}
