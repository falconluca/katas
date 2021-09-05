package concurrency

import (
	"fmt"
	"time"
)

func Select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "three"
	}()

	for i := 0; i < 3; i++ { // 因为需要获取三个chan的输出, 所以循环3次
		fmt.Printf("Select在这里阻塞住了, i:%v\n", i)
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func Timeouts() {
	c1 := make(chan string, 1)
	// external call
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
		fmt.Println("Note that the channel is buffered, so the send in the goroutine is nonblocking.")
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
