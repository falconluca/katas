package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ch := make(chan string, 1)

	for i := 0; i < 3; i++ {
		go func() {
			job := rand.Intn(10000)
			ch <- fmt.Sprintf("%d", job)
			fmt.Printf("生产者生产job: %d\n", job)
		}()
	}

	for i := 0; i < 3; i++ {
		time.Sleep(2 * time.Second)
		fmt.Printf("消费者消费第%v个job: %v\n\n", i+1, <-ch)
	}
	// Output:
	// 生产者生产job: 3125
	// 消费者消费第1个job: 3125
	//
	// 生产者生产job: 9580
	// 消费者消费第2个job: 9580
	//
	// 生产者生产job: 7576
	// 消费者消费第3个job: 7576
	//
	//
	// Other Output:
	// 生产者生产job: 1807
	// 消费者消费第1个job: 1807
	//
	// 生产者生产job: 6241
	// 生产者生产job: 8258
	// 消费者消费第2个job: 6241
	//
	// 消费者消费第3个job: 8258

	// 阻塞主进程
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
