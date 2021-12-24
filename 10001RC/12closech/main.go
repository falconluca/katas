package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "greetings"
		ch <- "luca"
		close(ch)
	}()

	var msg string
	unclose := true
	for unclose {
		select {
		case msg, unclose = <-ch:
			if unclose {
				fmt.Println(msg)
			} else {
				fmt.Println("channel has been closed!")
			}
		}
	}
}
