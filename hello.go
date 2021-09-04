package main

import (
	"awesomeProject/dao"
	"fmt"
	"time"
)

func main() {
	fmt.Println(dao.Insert())
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Weekday())
	fmt.Println(time.Saturday)
	fmt.Println(time.Sunday)
}
