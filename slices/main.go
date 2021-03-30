package main

import "fmt"

func main() {
	//A slice literal is declared just like an array literal, except you leave out the element count:
	var grades = []string{"A", "B", "C", "D", "Z"}
	grades = append(grades, "F")

	for i := 0; i < len(grades); i++ {
		fmt.Printf("Grade %d is %s\n", i, grades[i])
	}

	//lets copy a slice
	var uniqueGrades = make([]string, len(grades))
	copy(uniqueGrades, grades)

	fmt.Println(uniqueGrades)

	//use slice operator [low:high]
	var selectedGrades = grades[1:3]
	fmt.Println(selectedGrades)
}
