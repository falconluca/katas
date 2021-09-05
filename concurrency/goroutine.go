package concurrency

import (
	"fmt"
	"math/rand"
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

func Entry() {
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
