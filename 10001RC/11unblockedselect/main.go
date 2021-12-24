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
		default:
			fmt.Println("啥都没收到~")
			time.Sleep(time.Second)
		}
	}
	// Output:
	// 啥都没收到~
	// 接收到 luca
	// 啥都没收到~
	// 啥都没收到~
}
