package main

import "fmt"

type Article struct {
	id    int
	title string
}

// Go方法就是带接收者的函数

func (article Article) GetInfo() string {
	return fmt.Sprintf("id: %v, title: %s", article.id, article.title)
}

type Integer int

func (integer Integer) GetType() string {
	return fmt.Sprintf("typo: %T", integer)
}
