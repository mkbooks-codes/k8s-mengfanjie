package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	timer := time.NewTimer(time.Second)
	select {
	// check normal channel
	case <-ch:
		fmt.Println("received from ch")
	case <-timer.C:
		fmt.Println("timeout waiting from channel ch")
	}
}
