package concurrency

import "fmt"

func Channel() {
	// https://stackoverflow.com/questions/11943841/what-is-channel-buffer-size
	// buffer为2
	channel := make(chan string, 2)

	// goroutine向chan中发送消息, 与主线程通信
	go func() {
		channel <- "hello"
		channel <- "golang"
	}()

	// 阻塞等待
	greeting := <-channel
	name := <-channel
	fmt.Println(greeting)
	fmt.Println(name)
}
