package main

import (
	//"awesomeProject/dao"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Weekday())
	fmt.Println(time.Saturday)
	fmt.Println(time.Sunday)

	//dao.TestDb()
}
