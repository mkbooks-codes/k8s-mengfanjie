package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var c = make(chan string)
	go prod(c)
	go consume(c)
	time.Sleep(time.Second * 10)
}

func prod(sendOnly chan<- string) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				time.Sleep(time.Second)
				s := strconv.Itoa(i) + ": " + strconv.Itoa(j)
				sendOnly <- s
				fmt.Printf("prod-%d putting: %v \n", i, s)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func consume(readOnly <-chan string) {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			j := <-readOnly
			time.Sleep(time.Second)
			fmt.Printf("consume-%d receiving: %s \n", i, j)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
