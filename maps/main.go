package main

import "fmt"

func main() {
	//always use make to initialize a map
	var colours = make(map[string]string)

	colours["red"] = "#FF0000"
	colours["blue"] = "#0000FF"
	colours["yellow"] = "#FFFF00"

	fmt.Println("All colours: ", colours)
	fmt.Println("Colours map length: ", len(colours))
	fmt.Println("Value for red colour: ", colours["red"])

	delete(colours, "blue")

	_, found := colours["blue"]
	fmt.Println("\nBlue colour exists: ", found)

	//iterate through map
	for key, value := range colours {
		fmt.Println("Colour: ", key, "=>", "Code: ", value)
	}
}
