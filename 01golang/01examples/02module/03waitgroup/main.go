package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitBySleep()
	waitByChannel()
	waitByWG()
}

// 通过休眠保证所有协程执行完毕
func waitBySleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

// 通过 channel 保证所有协程执行完毕
func waitByChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}

// 通过 WaitGroup 保证所有协程执行完毕
func waitByWG() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
