package grammar

import (
	"fmt"
	"strings"
)

func GetPlayerNumber(name string) int {
	var result int
	switch playerName := "player: " + name; playerName {
	case "player: Luca":
		result = 0
		// Go 自动提供了 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止
	case "player: Allen":
		result = 1
	default:
		result = -1
	}
	return result
}

func GetPlayerNumberV2(name string) int {
	var result int
	playerName := "player: " + name
	switch {
	case strings.EqualFold(playerName, "player: Luca"):
		result = 0
		// Go 自动提供了 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止
	case strings.EqualFold(playerName, "player: Allen"):
		result = 1
	default:
		result = -1
	}
	return result
}

func V3(value interface{}) string {
	whatType := func(value interface{}) string {
		var result string
		switch typo := value.(type) {
		case bool:
			result = fmt.Sprintln("I am a bool.")
		case int:
			result = fmt.Sprintln("I am a int.")
		case string:
			result = fmt.Sprintln("I am a string.")
		default:
			result = fmt.Sprintf("Do not know typo %T.\n", typo)
		}
		return result
	}
	return whatType(value)
}
