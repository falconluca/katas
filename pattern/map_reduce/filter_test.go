package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	lst := []int{0, 2, 0, 4}
	res := IntFilter(lst, func(item int) bool {
		return item != 0
	})
	assert.Equal(t, 2, len(res))
	assert.Equal(t, 2, res[0])
	assert.Equal(t, 4, res[1])
}
