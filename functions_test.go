package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 函数一个返回值
func TestShouldReturnOneValue(t *testing.T) {
	assert.Equal(t, "Hey Luca", SayHey("Luca"))
}

// 函数多个返回值
func TestShouldReturnTwoValue(t *testing.T) {
	orderStatus, isSuccess := GetOrderStatusById(0, 0)
	assert.Equal(t, "unpaid", orderStatus)
	assert.True(t, isSuccess)

	orderStatus1, err1 := GetOrderStatusById(5, 0)
	// map的value类型为string默认值为空字符串
	assert.Equal(t, "", orderStatus1)
	assert.False(t, err1)

	first, second := SwapString("Hello", "Luca")
	assert.Equal(t, "Luca", first)
	assert.Equal(t, "Hello", second)
}

// 函数可变参数列表
func TestDynamicParams(t *testing.T) {
	totalGoodsStock := CalculateTotalGoodsStock(12, 17, 3)
	assert.Equal(t, 32, totalGoodsStock)

	goodsStockList := []int{20, 30, 12, 7}
	// Cannot use type '[3]int' as the type '[]int'
	//goodsStockList := [4]int{20, 30, 12, 7}
	assert.Equal(t, 69, CalculateTotalGoodsStock(goodsStockList...))
}

// TODO 函数闭包

// 函数递归
func TestRecursion(t *testing.T) {
	assert.Equal(t, 6, Fb(3))
	assert.Equal(t, 5040, Fb(7))
}
