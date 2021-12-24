package main

import "fmt"

func greetings(msg string) {
	fmt.Println(msg)
}

func main() {
	// go关键词类似C语言里的 pthread_create
	go func() {
		fmt.Println("hello")
	}()

	go greetings("hello Luca~")

	// 主进程退出之后 goroutine也就退出了, 因此控制台中看不到任何输出
}
