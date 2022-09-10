package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopFunc: ", i)
		}(i)
	}
}

// 会死锁，因为 defer 只会在方法退出时才执行，在方法退出前还在栈里，无法执行。通过匿名函数解决该问题。
// func loopFunc() {
// 	lock := sync.Mutex{}
// 	for i := 0; i < 3; i++ {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		fmt.Println("loopFunc: ", i)
// 	}
// }
