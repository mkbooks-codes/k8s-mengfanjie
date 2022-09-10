package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("hello from goroutine")
		ch <- 0 //数据写入Channel
	}()
	i := <-ch //从Channel中取数据并赋值
	fmt.Println(i)
}
