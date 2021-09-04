package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariable(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("https://www.baidu.com", GetConstString())
	assert.Equal("", DefaultValueOfString())
	assert.Equal("Allen", ValueOfString())

	name, _ := FormatPlayerName("Luca")
	assert.Equal("player name is Luca", name)

	_, err := FormatPlayerName("")
	assert.NotEmpty(err)
	assert.Equal("name can not be empty", err.Error())

	assert.Equal("player is Luca", ShortDeclare("Luca"))
}
