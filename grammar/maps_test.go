package grammar

import (
	"fmt"
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

	players := make(map[string]struct {
		id   string
		data float32
	})
	// TODO(不太明白) 映射的零值为nil
	assert.Empty(players)

	players["Luca"] = struct {
		id   string
		data float32
	}{
		"4068433", 74.39967,
	}
	assert.Equal(float32(74.39967), players["Luca"].data)
	assert.Equal("4068433", players["Luca"].id)

	goodPlayers := map[string]struct {
		id   string
		data float32
	}{
		"Luca": struct {
			id   string
			data float32
		}{
			"4068433", 74.39967,
		},
		"Allen": struct {
			id   string
			data float32
		}{
			"4068434", 84.39967,
		},
		// 若顶级类型只是一个类型名，你可以在文法的元素中省略它
		"Curry": {
			"4068436", 74.45967,
		},
	}
	assert.Equal(3, len(goodPlayers))

	player, exist := goodPlayers["Curry"]
	assert.Equal(true, exist)
	assert.Equal("4068436", player.id)

	player, exist = goodPlayers["Bruce"]
	assert.Equal(false, exist)
	// 若 key 不在映射中，那么 player 是该映射元素类型的零值
	assert.Equal("", player.id)
	assert.Equal(float32(0), player.data)
}

func TestMapKeySlices(t *testing.T) {
	m := map[string][]string{
		"all": {
			"hello",
			"luca",
			"James",
		},
		"all2": {
			"hello",
			"luca",
			"James",
		},
	}
	strings := m["all"]
	for _, str := range strings {
		fmt.Println(str)
	}
}
