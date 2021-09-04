package main

import (
	"github.com/stretchr/testify/assert"
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
