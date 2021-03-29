package main

import "fmt"

func main() {
	var grades = []string{"A", "B", "C", "D", "Z"}

	for i := 0; i < len(grades); i++ {
		fmt.Printf("Grade %d is %s\n", i, grades[i])
	}

	var twoDimensional [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoDimensional)
}
