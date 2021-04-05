package main

import (
	"fmt"
	"time"
)

func main() {
	//create a buffer with capacity of 2
	done := make(chan bool, 1)

	go doSomething(done)

	<-done
}

func doSomething(done chan bool) {
	fmt.Print("Working")
	time.Sleep(1 * time.Second)

	fmt.Print(".")
	time.Sleep(1 * time.Second)

	fmt.Print(".")
	time.Sleep(1 * time.Second)

	fmt.Print(".")
	time.Sleep(1 * time.Second)
	fmt.Println(" Done")

	done <- true
}
