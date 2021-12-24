package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 6)
		ch1 <- "hello"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- "luca"
	}()

	for i := 0; i < 4; i++ {
		select { // 可以使用select来监听两个ch, 还有一点就是select这个东西是阻塞的
		case msg1 := <-ch1:
			fmt.Println("接收到", msg1)
		case msg2 := <-ch2:
			fmt.Println("接收到", msg2)
		case <-time.After(time.Second * 3):
			fmt.Println("超时了...")
		}
	}
	// Output:
	// 接收到 luca
	// 超时了...
	// 接收到 hello
	// 超时了...
}
