package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	var lst = []string{"Luca", "Allen", "Curry"}
	res1 := MapToStr(lst, func(item string) string {
		return item + ","
	})
	assert.Equal(t, "Luca,", res1[0])
	assert.Equal(t, "Allen,", res1[1])
	assert.Equal(t, "Curry,", res1[2])

	res2 := MapToInt(lst, func(item string) int {
		if item == "Luca" {
			return 1
		} else {
			return 0
		}
	})
	assert.Equal(t, 1, res2[0])
	assert.Equal(t, 0, res2[1])
	assert.Equal(t, 0, res2[2])
}
