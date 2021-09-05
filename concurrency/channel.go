package concurrency

import "fmt"

func Channel() {
	channel := make(chan string)

	// goroutine向chan中发送消息, 与主线程通信
	go func() {
		channel <- "hello"
	}()

	// 阻塞等待
	result := <-channel
	fmt.Println(result)
}
