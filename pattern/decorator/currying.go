package decorator

import "fmt"

func decorator(fn func(string)) func(string) {
	return func(str string) {
		fmt.Println("Starting...")
		fn(str)
		fmt.Println("Ending...")
	}
}

func Hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}

// https://zh.wikipedia.org/wiki/%E6%9F%AF%E9%87%8C%E5%8C%96
func Curry() {
	xHello := decorator(Hello)
	xHello("Luca")

	decorator(Hello)("Curry")
}
