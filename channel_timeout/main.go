package main

import "time"

func main() {
	channelOne := make(chan string, 1)

	//Immediately invoked function expression (IIFE)
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "result one"
	}()
}
