package main

import (
	"strings"
)

func CommonFor() string {
	var result strings.Builder

	players := map[int]string{0: "Luca", 1: "Allen", 2: "Curry", 3: "James"}

	result.WriteString("players: ")
	for i := 0; i < len(players); i++ {
		playerName := players[i]

		result.WriteString(playerName)
		result.WriteString(", ")
	}
	s := result.String()
	return s[0 : len(s)-2] // 移除最后一个逗号和空格
}

func WhileLikeFor() string {
	var result strings.Builder

	players := map[int]string{0: "Luca", 1: "Allen", 2: "Curry", 3: "James"}

	result.WriteString("players: ")
	i := 0
	for i < len(players) {
		playerName := players[i]

		result.WriteString(playerName)
		result.WriteString(", ")

		if players[i] != "" {
			i++
		}
	}
	s := result.String()
	return s[0 : len(s)-2]
}

func InfiniteFor() string {
	var result strings.Builder

	players := map[int]string{0: "Luca", 1: "Allen", 2: "Curry", 3: "James"}

	result.WriteString("players: ")
	i := 0
	for {
		playerName := players[i]
		if playerName == "" {
			break
		}
		result.WriteString(playerName)
		result.WriteString(", ")

		i++
	}
	s := result.String()
	return s[0 : len(s)-2]
}
