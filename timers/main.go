package main

import (
	"fmt"
	"time"
)

//TODO: practise more
func main() {
	timerOne := time.NewTimer(3 * time.Second)

	<-timerOne.C
	fmt.Println("Timer expired")
}
