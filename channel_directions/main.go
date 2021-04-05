package main

import "fmt"

func ping(pings chan<- string, message string) {
	pings <- message
}

//Pay attention to the position of arrows for the 2 arguements
func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	pongs <- message
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Hello there")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
