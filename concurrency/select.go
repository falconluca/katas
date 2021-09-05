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
