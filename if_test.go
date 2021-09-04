package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
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

func Conditions(idx int) string {
	if idx < 0 || idx > 2 {
		return ""
	}

	players := [5]string{"Luca", "Allen", "CURRY"}
	player := players[idx]
	if strings.EqualFold(player, "LUCA") {
		return "is Luca"
	} else if player == "Allen" {
		return "is Allen"
	} else {
		return "is Curry"
	}
}
