package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan int)
	go prod(c)
	go consume(c)
	time.Sleep(time.Second * 10)
}

func prod(sendOnly chan<- int) {
	for {
		time.Sleep(time.Second)
		sendOnly <- 1
		fmt.Println("putting: ", 1)
	}
}

func consume(readOnly <-chan int) {
	for {
		time.Sleep(time.Second)
		i := <-readOnly
		fmt.Println("receiving: ", i)
	}
}
