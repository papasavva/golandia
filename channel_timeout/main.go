package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne := make(chan string, 1)

	//Immediately invoked function expression (IIFE)
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "result one"
	}()

	select {
	case response := <-channelOne:
		fmt.Println(response)
	case <-time.After(1 * time.Second):
		fmt.Println("First timeout")
	}

	channelTwo := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		channelTwo <- "result two"
	}()

	select {
	case response := <-channelTwo:
		fmt.Println(response)
	case <-time.After(3 * time.Second):
		fmt.Println("Second timeout")
	}
}
