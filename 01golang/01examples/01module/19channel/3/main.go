package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10) // n will be between 0 and 10
			fmt.Println("putting: ", n)
			ch <- n
		}
	}()

	if v, notClosed := <-ch; notClosed {
		fmt.Println("receiving: ", v)
	}

}
