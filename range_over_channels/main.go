package main

import "fmt"

func main() {
	c := make(chan string, 4)
	c <- "hi"
	c <- "there"

	close(c)

	for element := range c {
		fmt.Println(element)
	}

}
