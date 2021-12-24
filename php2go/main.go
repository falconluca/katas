package main

import (
	"fmt"
	"github.com/syyongx/php2go"
)

func main() {
	ExampleDateTime()
}

func ExampleDateTime() {
	fmt.Println(php2go.Time()) // 获取当前时间戳
	// Output:
	// 1638253730

	// 字符串日期转时间戳
	strtotime, _ := php2go.Strtotime("2006-01-02 15:04:05", "2021-11-30 14:28:50")
	fmt.Println(strtotime)
	// Output:
	// 1638282530

	// 时间戳转字符串日期
	date := php2go.Date("2006-01-02 15:04:05", 1638253730)
	fmt.Println(date)
	// Output:
	// 2021-11-30 14:28:50

	fmt.Println(php2go.Checkdate(10, 32, 2021)) // 检查是否为合法的公历日期
	// Output:
	// false
	fmt.Println(php2go.Checkdate(10, 30, 2021))
	// Output:
	// true
}

func ExampleString() {
	result := php2go.Strpos("happy halloween!", "halloween", 1)
	fmt.Println(result)
	// Output:
	// 6

	// TODO
}

func ExampleArray() {
	// TODO
}

func ExampleNetwork() {
	fmt.Println(php2go.Gethostname())

	ip2long := php2go.IP2long("192.168.10.1")
	fmt.Println(ip2long)
	// Output:
	// 3232238081

	long2ip := php2go.Long2ip(3232238081)
	fmt.Println(long2ip)
	// Output:
	// 3232238081
}

func ExampleMisc() {
	_ = php2go.Putenv("firstName=Luca") // TODO sync.Once
	firstName := php2go.Getenv("firstName")
	fmt.Println(firstName)
	// Output:
	// Luca

	usage := php2go.MemoryGetUsage(true)
	fmt.Println(usage)
}

func ExampleSleep() {
	// 休眠3秒
	//php2go.Sleep(3)

	// 休眠1秒
	//php2go.Usleep(1000000) // 1微秒等于一百万分之一秒
}
