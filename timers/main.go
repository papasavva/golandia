package main

import (
	"fmt"
	"time"
)

func main() {
	timerOne := time.NewTimer(3 * time.Second)

	<-timerOne.C
	fmt.Println("Timer expired")
}
