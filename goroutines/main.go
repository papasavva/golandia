package main

import (
	"fmt"
	"time"
)

func printXtimes(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Text:", text, "Itteration", i)
	}
}
func main() {
	//two function calls are running asynchronously
	go printXtimes("Hello", 100)
	go printXtimes("there", 100)

	time.Sleep(3 * time.Second)

}
