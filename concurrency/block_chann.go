package concurrency

import "fmt"
import "time"

func BlockChannel() {
	// buffer为1
	channel := make(chan string)

	go func() {
		channel <- "hello" // 阻塞
		fmt.Println("write \"hello\" done!")

		channel <- "World" // Reader在Sleep，这里在阻塞
		fmt.Println("write \"World\" done!")

		fmt.Println("Write go sleep...")
		time.Sleep(3 * time.Second)
		channel <- "channel"
		fmt.Println("write \"channel\" done!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Reader Wake up...")
	msg := <-channel

	fmt.Println("Reader: ", msg)
	msg = <-channel

	fmt.Println("Reader: ", msg)

	msg = <-channel // Writer在Sleep，这里在阻塞
	fmt.Println("Reader: ", msg)
}

func worker(done chan bool) {
	fmt.Println("Worker is working...")

	time.Sleep(time.Second)

	fmt.Println("Done!")
	done <- true
}

func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done)

	// 利用channel阻塞的特性实现主进程等待的效果
	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
	//<- pings // error!
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "Let's play Overwatch together!")
	pong(pings, pongs)

	result := <-pongs
	fmt.Println(result)
}
