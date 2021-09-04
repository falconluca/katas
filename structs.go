package main

import "strconv"

type TiktokNews struct {
	id    int64
	title string
	url   string
}

// GO没有访问控制符号，如果你想让一个方法可以被别的包访问的话，你需要把这个方法的第一个字母大写。这是一种约定。

func (tiktokNews *TiktokNews) Parse() string {
	return "id: " + strconv.FormatInt(tiktokNews.id, 10) +
		", title: " + tiktokNews.title + ", url: " + tiktokNews.url
}

func (tiktokNews *TiktokNews) ParseWithBr(br string) string {
	return "id: " + strconv.FormatInt(tiktokNews.id, 10) +
		br + "title: " + tiktokNews.title +
		br + "url: " + tiktokNews.url
}
