package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetArticleInfo(t *testing.T) {
	article := Article{id: 1000, title: "北京证劵交易所"}
	assert.Equal(t, "id: 1000, title: 北京证劵交易所", article.GetInfo())

	phoneNumber := Integer(13000000000)
	assert.Equal(t, "typo: main.Integer", phoneNumber.GetType())

	post := Post{Id: 10001, Title: "一片短文"}
	assert.Equal(t, "id: 10001, title: 一片短文", post.GetInfo())
	post.UpdatePost("一篇短文")
	assert.Equal(t, "id: 10001, title: 一篇短文", post.GetInfo())
}
