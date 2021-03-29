package main

import "fmt"

func main() {
	var grades = []string{"A", "B", "C", "D", "Z"}

	for i := 0; i < len(grades); i++ {
		fmt.Printf("Grade %d is %s\n", i, grades[i])
	}

}
