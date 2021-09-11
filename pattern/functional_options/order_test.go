package functional_options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrder(t *testing.T) {
	assert2 := assert.New(t)

	orderStatus := OrderStatus{code: 2, msg: "已完成"}
	order, _ := NewOrder("OrderId", Status(orderStatus), GoodsId("GoodsId"))
	assert2.Equal("已完成", order.Status.msg)
	assert2.Equal(2, order.Status.code)
	assert2.Equal("OrderId", order.OrderId)
}
