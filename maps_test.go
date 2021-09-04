package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaps(t *testing.T) {
	assert := assert.New(t)

	orderStatus := make(map[string]int)
	orderStatus["unpaid"] = 0
	orderStatus["paid"] = 2
	orderStatus["finished"] = 3

	assert.Equal(3, len(orderStatus))

	for k, value := range orderStatus {
		assert.Equal(value, orderStatus[k])
	}

	paidValue := orderStatus["paid"]
	assert.Equal(2, paidValue)

	// Expected nil, but got: 0
	//assert.Nil(orderStatus["nil"])
	assert.Empty(orderStatus["nil"])

	delete(orderStatus, "finished")
	assert.Equal(2, len(orderStatus))

	netIds := map[string]int{"threeNet": 0, "ftx": 1, "xfl": 2}
	assert.Equal(3, len(netIds))
	assert.Equal(0, netIds["threeNet"])
	assert.Equal(1, netIds["ftx"])
	assert.Equal(2, netIds["xfl"])

	delete(netIds, "threeNet")
	delete(netIds, "ftx")
	delete(netIds, "xfl")
	assert.Empty(netIds)
}
