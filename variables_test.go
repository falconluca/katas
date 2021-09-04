package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVariable(t *testing.T) {
	assert := assert.New(t)

	url, enable := GetConst()
	assert.Equal("https://www.baidu.com", url)
	assert.Equal(true, enable)

	playerName, playerNumber, isStarting := DefaultValues()
	assert.Equal("", playerName)
	assert.Equal(0, playerNumber)
	assert.Equal(false, isStarting)

	assert.Equal("Allen", ValueOfString())

	name, _ := FormatPlayerName("Luca")
	assert.Equal("player name is Luca", name)

	_, err := FormatPlayerName("")
	assert.NotEmpty(err)
	assert.Equal("name can not be empty", err.Error())

	assert.Equal("player is Luca", ShortDeclare("Luca"))
}
