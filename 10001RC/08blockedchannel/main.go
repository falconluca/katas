package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("生产者开始生产第一个job: hello")
		ch <- "hello" // 阻塞在这
		fmt.Println("生产者生产第一个job生产完毕")

		fmt.Println("生产者生产第二个job: luca")
		ch <- "luca"
		fmt.Println("生产者生产第二个job生产完毕")

		fmt.Println("Oops! 生产者被挂起")
		time.Sleep(3 * time.Second)
		fmt.Println("生产者结束挂起")

		fmt.Println("生产者生产第三个job: !")
		ch <- "!"
		fmt.Println("生产者生产第三个job生产完毕")
	}()

	fmt.Println("Oops! 消费者被挂起")
	time.Sleep(3 * time.Second)
	fmt.Println("消费者结束挂起")

	job := <-ch
	fmt.Println("消费者消费第一个job", job)

	job = <-ch
	fmt.Println("消费者消费第二个job", job)

	job = <-ch
	fmt.Println("消费者消费第三个job", job)
}
