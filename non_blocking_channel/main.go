package main

import "fmt"

func main() {
	messages := make(chan string)

	//receive
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received") //no value available in messages, so this will get executed
	}

	//send
	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("Sent message:", msg)
	default:
		fmt.Println("No message sent") //no buffer and no receiver, so this will be executed
	}
}
