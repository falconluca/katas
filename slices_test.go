package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
// see the example in the comments areas.
func TestSlices(t *testing.T) {
	assert := assert.New(t)

	// 数组的长度是其类型的一部分，因此数组不能改变大小。
	players := [5]string{"Luca", "Allen", "Curry", "James", "Love"} // 数组
	assert.Equal(5, len(players))
	assert.Equal("Luca", players[0])
	assert.Equal("Curry", players[2])
	assert.Equal("Love", players[4])

	// 切片就像数组的引用
	var nbaPlayers []string = players[1:] // 切片
	assert.Equal(4, len(nbaPlayers))
	assert.Equal("Allen", nbaPlayers[0])
	assert.Equal("Love", nbaPlayers[3])

	// Golang切片是共享内存的，没有数据的复制，只是记录从哪切到哪的信息
	players[2] = "CURRY"
	assert.Equal("CURRY", nbaPlayers[1])

	var goodPlayers []string = players[:3] // 切片
	assert.Equal(3, len(goodPlayers))
	assert.Equal("Luca", goodPlayers[0])
	assert.Equal("Allen", goodPlayers[1])
	assert.Equal("CURRY", goodPlayers[2])

	opponents := players[2:4] // 切片
	assert.Equal(2, len(opponents))
	assert.Equal("CURRY", opponents[0])
	assert.Equal("James", opponents[1])

	// 数组默认值
	var buckets [3]int
	assert.Equal(len(buckets), 3)
	assert.Equal(0, buckets[0])
	assert.Equal(0, buckets[1])
	assert.Equal(0, buckets[2])

	buckets[2] = 7
	assert.Equal(0, buckets[1])
	assert.Equal(7, buckets[2])
}

func TestDifferenceBetweenArrayAndSlice(t *testing.T) {
	// 创建一个数组
	var players [3]string = [3]string{"Luca", "Allen", "James"}
	assert.Equal(t, "[3]string", reflect.TypeOf(players).String())

	// 通过数组创建一个切片
	var playersSlice []string = players[0:]
	assert.Equal(t, "[]string", fmt.Sprintf("%T", playersSlice))

	// 直接创建一个切片
	// 工作原理：先创建数组，然后基于原数组构建一个引用了它的切片
	var language []string = []string{"Go", "Java", "Bash"}
	assert.Equal(t, "[]string", reflect.TypeOf(language).String())

	// 创建结构体切片
	var tiktokNews []TiktokNews = []TiktokNews{
		{id: 1000, title: "早上吃什么", url: "https://www.baidu.com"},
		{id: 1001, title: "中午上吃什么", url: "https://www.douban.com"},
		{id: 1002, title: "晚上吃什么", url: "https://www.zhihu.com"},
	}
	assert.Equal(t, "[]main.TiktokNews", reflect.TypeOf(tiktokNews).String())

	// 匿名结构体切片
	var books []struct {
		id   int
		name string
	} = []struct {
		id   int
		name string
	}{
		{0, "领域驱动设计"},
		{1, "代码整洁之道"},
		{2, "重构"},
	}
	assert.Equal(t, "[]struct { id int; name string }", reflect.TypeOf(books).String())

	// 切片的切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	assert.Equal(t, "X", board[0][0])
	assert.Equal(t, "O", board[2][2])
	assert.Equal(t, "X", board[1][2])
	assert.Equal(t, "O", board[1][0])
	assert.Equal(t, "X", board[0][2])
}

func TestDefaultBehaviorOfSlices(t *testing.T) {
	var players [10]string
	// 数组和切片不是同一种数据类型
	assert.NotEqual(t, players, players[0:10])
	assert.NotEqual(t, players, players[:10])
	assert.NotEqual(t, players, players[0:])
	assert.NotEqual(t, players, players[:])

	playersSlice := players[0:10]
	assert.Equal(t, playersSlice, players[0:])
	assert.Equal(t, playersSlice, players[:10])
	assert.Equal(t, playersSlice, players[:])
}

func TestSliceCapAndLen(t *testing.T) {
	players := []string{"Luca", "Allen", "Curry", "James", "Love", "Yao"}
	assert.Equal(t, 6, len(players))
	assert.Equal(t, 6, cap(players))

	// len: 切片的长度就是它所包含的元素个数
	assert.Equal(t, 4, len(players[2:]))
	// cap: 切片的容量是从它的第一个元素开始数，到其【底层数组】元素末尾的个数
	assert.Equal(t, 4, cap(players[2:]))

	assert.Equal(t, 2, len(players[:2]))
	assert.Equal(t, 6, cap(players[:2]))

	assert.Equal(t, 2, len(players[2:4]))
	assert.Equal(t, 4, cap(players[2:4]))
}

func TestDefaultValueOfSlice(t *testing.T) {
	var players []string
	// nil 切片的长度和容量为 0 且没有底层数组
	assert.Equal(t, 0, len(players))
	assert.Equal(t, 0, cap(players))
	assert.Equal(t, []string(nil), players)

	var phoneNumbers []int
	assert.Equal(t, 0, len(phoneNumbers))
	assert.Equal(t, 0, cap(phoneNumbers))
	assert.Equal(t, []int(nil), phoneNumbers)

	var isStarting []bool
	assert.Equal(t, 0, len(isStarting))
	assert.Equal(t, 0, cap(isStarting))
	assert.Equal(t, []bool(nil), isStarting)
}
