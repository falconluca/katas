package concurrency

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func routine(name string, delay time.Duration) {
	t0 := time.Now()
	fmt.Println(name, " 开始处理, 开始时间: ", t0)

	time.Sleep(delay)

	t1 := time.Now()
	fmt.Println(name, " 处理完毕, 结束时间: ", t1)
	fmt.Println(name, " 处理时间: ", t1.Sub(t0))
}

func GetStarted() {
	// 生成随机种子
	rand.Seed(time.Now().Unix())

	var name string
	for i := 0; i < 3; i++ {
		name = fmt.Sprintf("Worker_%02d", i)

		// 生成随机等待时间，从0-4秒
		delay := time.Duration(rand.Intn(5)) * time.Second

		go routine(name, delay)
	}

	// 使用go关键字来调用匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("Hello goroutine!")

	// 阻塞主进程停住，否则主进程退了，goroutine也就退了
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

var totalTickets int32 = 10
var mutex sync.Mutex

func SellTickets(i int) {
	for totalTickets > 0 {
		mutex.Lock()
		if totalTickets > 0 { // 如果有票就卖
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond) // 休息一下

			// 卖一张票
			totalTickets--

			fmt.Println("买票人id:", i, ", 买了一张票, 剩余票数:", totalTickets)
		}
		mutex.Unlock()
	}
}

func dispatchSellTickets() {
	runtime.GOMAXPROCS(8) // 我的电脑是8核处理器，所以我设置了8
	rand.Seed(time.Now().Unix())

	// 并发5个goroutine来卖票
	for i := 0; i < 5; i++ {
		go SellTickets(i)
	}

	var input string
	fmt.Scanln(&input)
	// 退出时打印还有多少票
	fmt.Printf("剩余票数: %v\n", totalTickets)
	fmt.Println("done")
}

func Cas() {
	var counter uint32 = 0
	for i := 0; i < 10; i++ {
		// 开启10个goroutine
		go func() {
			for i := 0; i < 20; i++ {
				time.Sleep(time.Millisecond)
				// https://pkg.go.dev/sync/atomic
				atomic.AddUint32(&counter, 1)
			}
		}()
	}

	// 休眠一秒钟等goroutine完成
	time.Sleep(time.Second)

	result := atomic.LoadUint32(&counter)
	fmt.Println("counter:", result)
}

func Entry() {
	//GetStarted()
	//dispatchSellTickets()
	//Cas()
	//Channel()
	//BlockChannel()
	//ChannelSync()
	ChannelDirections()
}
