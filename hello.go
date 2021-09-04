package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//
	// for
	for page := 1; page <= 10; page++ {
		fmt.Printf("Current page: %v \n", page)
	}

	//
	// while-like-for
	page := 1
	for page <= 10 {
		fmt.Printf("Current page: %v \n", page)
		page++
	}

	//
	// if
	username := ""
	if len(username) == 0 { // 条件表达式不需要携带()
		fmt.Println("username can not blank.")
	}

	//
	// if-elif-else
	if username := "luca"; strings.EqualFold(username, "Luca") {
		fmt.Println("Welcome back. Luca.")
	} else if strings.EqualFold(username, "curry") {
		fmt.Println("Oh is Curry!")
	} else {
		fmt.Printf("Unknown username: %s \n", username)
	}

	//
	// switch by typo of input
	whatType := func(value interface{}) {
		switch typo := value.(type) {
		case bool:
			fmt.Println("I am a bool.")
		case int:
			fmt.Println("I am a int.")
		case string:
			fmt.Println("I am a string.")
		default:
			fmt.Printf("Do not know typo %T\n", typo)
		}
	}
	whatType(12)
	whatType("luca")
	whatType(false)
	whatType(13.14)

	//
	// switch
	name := "james"
	switch name {
	case "luca":
		fmt.Println("hello luca!")
	case "allen", "james":
		fmt.Println("hello allen or james?")
	default:
		fmt.Println("what is your name?")
	}

	//
	// switch case语句支持表达式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//
	// time go builtin api
	fmt.Println(t)
	fmt.Println(t.Weekday())
	fmt.Println(time.Saturday)
	fmt.Println(time.Sunday)
}
