package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStruct(t *testing.T) {
	assert := assert.New(t)

	tiktokNews := TiktokNews{id: 10000, title: "早餐吃啥呢", url: "https://www.baidu.com"}
	assert.Equal(int64(10000), tiktokNews.id)
	assert.Equal("早餐吃啥呢", tiktokNews.title)
	assert.Equal("https://www.baidu.com", tiktokNews.url)

	// 变量调用结构体方法
	assert.Equal("id: 10000, title: 早餐吃啥呢, url: https://www.baidu.com", tiktokNews.Parse())
	expected := "id: 10000\ntitle: 早餐吃啥呢\nurl: https://www.baidu.com"
	assert.Equal(expected, tiktokNews.ParseWithBr("\n"))
	expected = "id: 10000\ttitle: 早餐吃啥呢\turl: https://www.baidu.com"
	assert.Equal(expected, tiktokNews.ParseWithBr("\t"))

	// 获取结构体的指针
	pTiktokNews := &tiktokNews // &{10000 早餐吃啥呢 https://www.baidu.com}
	(*pTiktokNews).title = "吃点好的"
	(*pTiktokNews).id = 10001
	assert.Equal(int64(10001), (*pTiktokNews).id)
	assert.Equal("吃点好的", (*pTiktokNews).title)
	assert.Equal("https://www.baidu.com", (*pTiktokNews).url)

	// 指针调用结构体方法
	assert.Equal("id: 10001, title: 吃点好的, url: https://www.baidu.com", pTiktokNews.Parse())
	expected = "id: 10001\ntitle: 吃点好的\nurl: https://www.baidu.com"
	assert.Equal(expected, pTiktokNews.ParseWithBr("\n"))
	expected = "id: 10001\ttitle: 吃点好的\turl: https://www.baidu.com"
	assert.Equal(expected, pTiktokNews.ParseWithBr("\t"))
}
