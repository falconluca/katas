package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		ch <- "greetings!"
		fmt.Println("好的, 第一个数据写进去了!")
		ch <- "Luca"
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
