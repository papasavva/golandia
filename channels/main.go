package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go greet(message)

	message <- "Bob Sponge"

	time.Sleep(1 * time.Second)
	close(message)
}

func greet(c chan string) {
	//the function is blocked and waits to receive a value
	name := <-c
	fmt.Println("Hello", name)
}
