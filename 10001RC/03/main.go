package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// time.Duration 和 time.Time 的区别?

func job(name string, delay time.Duration) {
	beginAt := time.Now()
	fmt.Println(name, " start at ", beginAt)

	time.Sleep(delay)

	endAt := time.Now()
	fmt.Println(name, " end at ", endAt)

	fmt.Println(name, " lasted ", endAt.Sub(beginAt))
}

func main() {
	runtime.GOMAXPROCS(4)
	rand.Seed(time.Now().Unix())

	var name string
	for i := 0; i < 3; i++ {
		name = fmt.Sprintf("go_%02d", i)
		// 随机时间0-4秒
		du := time.Duration(rand.Intn(5)) * time.Second
		go job(name, du)
	}

	// 阻塞主进程
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
