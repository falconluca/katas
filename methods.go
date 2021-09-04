package main

import "fmt"

type Article struct {
	id    int
	title string
}

// Go方法就是带接收者的函数

func (article Article) GetInfo() string {
	// 这里把article的id修改为12, 只能影响这个方法的返回值, 不能影响到接收者的值.
	// 需要修改接收者的值, 请使用指针接收者.
	//article.id = 12
	return fmt.Sprintf("id: %v, title: %s", article.id, article.title)
}

type Integer int

func (integer Integer) GetType() string {
	return fmt.Sprintf("typo: %T", integer)
}

type Post struct {
	Id    int
	Title string
}

func (post *Post) UpdatePost(name string) {
	if name == "" {
		return
	}

	post.Title = name
}

func (post *Post) GetInfo() string {
	return fmt.Sprintf("id: %v, title: %s", post.Id, post.Title)
}
