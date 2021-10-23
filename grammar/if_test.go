package grammar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditions(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("is Luca", Conditions(0))
	assert.Equal("is Allen", Conditions(1))
	assert.Equal("is Curry", Conditions(2))
	assert.Empty(Conditions(-1))
	assert.Empty(Conditions(3))
}
