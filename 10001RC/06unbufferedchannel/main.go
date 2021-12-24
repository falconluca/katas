package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine 与 main(程序的第一个goroutine) 通信

	// 没有缓存的channel 请与07bufferedchannel进行比对
	//
	// 要知道 make(chan string) 和 make(chan string, 1) 这两个是不相同的
	// 第一个ch没有任何缓冲区 但是第二个ch是有一个缓冲区的, 通过这个程序可以调整出来(调整第21行代码观察结果有啥不同)
	ch := make(chan string, 1)

	go func() {
		ch <- "greetings!" // 如果没有对应准备好结束数据的goroutine, 程序执行到这里会被阻塞(调整第21行代码观察结果有啥不同)
		fmt.Println("好的, 第一个数据写进去了!")
		ch <- "Luca" // 程序执行的这里会被阻塞, 因为channel是没有缓存的, 因此只有当channel中的数据被消费之后, 这里的数据才会被写进去
		fmt.Println("行, 第二个数据也写进去了!!!")
	}()

	time.Sleep(3 * time.Second)

	msg := <-ch
	fmt.Println(msg)

	time.Sleep(3 * time.Second)

	msg2 := <-ch
	fmt.Println(msg2)

	// 阻塞主进程
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
