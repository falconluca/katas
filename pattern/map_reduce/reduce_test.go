package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce(t *testing.T) {
	lst := []string{"Luca", "Allen", "Curry"}
	// 统计字符个数
	result := Reduce(lst, func(item string) int {
		return len(item)
	})
	assert.Equal(t, 14, result)
}
