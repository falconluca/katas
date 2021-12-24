package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter uint32 = 0
	for i := 0; i < 10; i++ {
		go func() {
			// 每个goroutine对counter累加20次
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				atomic.AddUint32(&counter, 1)
			}
		}()
	}

	time.Sleep(time.Second)
	result := atomic.LoadUint32(&counter)
	fmt.Println("result: ", result)
}
