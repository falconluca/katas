package main

import "fmt"

func greetings(msg string) {
	fmt.Println(msg)
}

func main() {
	go func() {
		fmt.Println("hello")
	}()

	go greetings("hello Luca~")

	// 阻塞主进程
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
