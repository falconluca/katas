package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type (
	Result string
	Search func(query string) Result
)

var (
	// 模拟两台web服务实例
	Web  = fakeSearch("web")
	Web2 = fakeSearch("web2")

	// 模拟两台image服务实例
	Image  = fakeSearch("image")
	Image2 = fakeSearch("image2")

	// 模拟两台video服务实例
	Video  = fakeSearch("video")
	Video2 = fakeSearch("video2")
)

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

// Google 低配版 串行化执行...
func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// Google2 这是一个并行执行的程序, 但是我们这里没有使用到条件变量、锁、回调函数
func Google2(query string) (results []Result) {
	ch := make(chan Result)
	go func() {
		ch <- Web(query)
	}()
	go func() {
		ch <- Image(query)
	}()
	go func() {
		ch <- Video(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-ch
		results = append(results, result)
	}
	return
}

// Google2dot1 不等待慢的服务(img, v, web)响应 自动超时
func Google2dot1(query string) (results []Result) {
	ch := make(chan Result)
	go func() {
		ch <- Web(query)
	}()
	go func() {
		ch <- Image(query)
	}()
	go func() {
		ch <- Video(query)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	ch := make(chan Result)
	searchReplica := func(i int) {
		ch <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
		/*
			go func(i int) {
				ch <- replicas[i](query)
			}(i)
		*/
	}
	return <-ch
}

func Google3(query string) (results []Result) {
	ch := make(chan Result)
	go func() {
		ch <- First(query, Web, Web2)
	}()
	go func() {
		ch <- First(query, Image, Image2)
	}()
	go func() {
		ch <- First(query, Video, Video2)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

func main() {
	runtime.GOMAXPROCS(8)
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	//results := Google("golang")
	//results := Google2("golang")
	//results := Google2dot1("golang")
	results := Google3("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

/*
1. Concurrency(并发) is not parallelism(并行)

如果你的电脑只有一个核心, 那么你可以并发执行, 但是没有办法并行执行

2. go的一个特性： main函数退出之后，程序就退出了

3. 程序中可能只有一个线程，但是有上千万上goroutine，因此goroutine不是线程，而是独立，彼此互不相关的函数

4. sync

`<-c` 如果消费者没有准备发送数据时 或者 生产者没有准备生产数据时, 这个操作会被阻塞

同样

`c<-` 如果消费者没有准备消费数据时, 这个操作会被阻塞

**总之就是要双端都准备好，他们才能正常进行交流**

buffered removes sync同步 要特别注意!!! 会让ch发生异步的行为

goroutine之间的通信，同步

**go不是通过共享内存来实现通信的，而是通过通信来实现共享内存的** 因为这是同步的
===> CSP 这样就不用加锁、条件变量、信号等哪些复杂的东西了 `ch := make(chan string)` 就是这个东西哈哈哈

5. concurrency pattern

使用generator pattern 实现 multiplexing

7. 全局timeout or 针对每个消息的timeout

8. receive on quit chan

9. daisy-chain
*/
