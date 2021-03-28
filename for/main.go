package main

import "fmt"

func simpleLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println("index: ", i)
	}
}

func loopOverArray() {
	var grades = []uint{20, 19, 10}

	for i := 0; i < len(grades); i++ {
		//if grades[i] == 20 {
		//	continue
		//}
		fmt.Println("Grade: ", grades[i])
		//break
	}
}

func stringIteration() {
	const city string = "Βαβυλώνα"

	for index, character := range city {
		fmt.Printf("Character %c is at position %d\n", character, index)
	}

}
func main() {
	simpleLoop()
	loopOverArray()
	stringIteration()
}
