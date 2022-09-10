package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go prod(c)
	go consume(c)
	time.Sleep(time.Millisecond)
}

func prod(sendOnly chan<- int) {
	for {
		sendOnly <- 1
	}
}

func consume(readOnly <-chan int) {
	for {
		i := <-readOnly
		fmt.Println(i)
	}
}
