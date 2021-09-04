package main

import "strings"

func Conditions(idx int) string {
	if idx < 0 || idx > 2 {
		return ""
	}

	players := [5]string{"Luca", "Allen", "CURRY"}

	// 简短语句的player变量, 其作用域仅限于if结构内
	if player := players[idx]; strings.EqualFold(player, "LUCA") {
		return "is Luca"
	} else if player == "Allen" {
		return "is Allen"
	} else {
		return "is Curry"
	}
}
