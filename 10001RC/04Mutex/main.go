package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	totalTickets int32 = 10
	mutex        sync.Mutex
)

// 买票程序
func main() {
	// goroutine 是用操作系统的线程来实现的,
	// 但是 goroutine 有个特性: 如果一个goroutine没有被阻塞,
	// 那么别的 goroutine 就不会被执行到,
	// 类似一个线程被分解成多个 goroutine 因此也叫"轻量级线程" 哈哈哈
	// 这不是真正的并发
	// 因此, 如果我们想要真正的并发, 那就需要在main函数的第一行上添加 GOMAXPROCS 来控制调度器使用的CPU核心数
	runtime.GOMAXPROCS(8)

	rand.Seed(time.Now().Unix())

	for i := 0; i < 5; i++ {
		go sellTickets(i)
	}

	// 阻塞主进程
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func sellTickets(i int) {
	for {
		if totalTickets <= 0 {
			break
		}

		mutex.Lock()
		if totalTickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			totalTickets--
			fmt.Println("id: ", i, "got a ticket. total ticket: ", totalTickets)
		}
		mutex.Unlock()
	}
}
