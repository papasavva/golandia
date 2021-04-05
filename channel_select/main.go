package main

import (
	"fmt"
	"time"
)

func serverOne(channel chan string) {
	time.Sleep(5 * time.Second)
	channel <- "message from server 1"
}

func serverTwo(channel chan string) {
	time.Sleep(3 * time.Second)
	channel <- "message from server 2"
}

func main() {
	outputOne := make(chan string)
	outputTwo := make(chan string)

	go serverOne(outputOne)
	go serverTwo(outputTwo)

	select { //blocks until one of its cases is ready
	case s1 := <-outputOne:
		fmt.Println(s1)
	case s2 := <-outputTwo:
		fmt.Println(s2)

	}
}
