package main

import "fmt"

func main() {
	c := make(chan string, 4)
	c <- "hi"
	c <- "there"
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)
	//we cannot send data after channel is closed, it will throw an error
	//c <- "!"

}
