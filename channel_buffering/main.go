package main

import (
	"fmt"
	"time"
)

func main() {
	//create a buffer with capacity of 2
	message := make(chan string, 2)

	message <- "Bob Sponge"
	message <- "Patricio"
	//we send to a buffered channel block when the buffer is full

	go greet(message)
	go greet(message)
	time.Sleep(1 * time.Second)
	close(message)
}

func greet(c chan string) {
	//the function is blocked and waits to receive a value
	name := <-c
	fmt.Println("Hello", name)
}
