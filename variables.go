package main

import (
	"errors"
	"fmt"
)

func GetConstString() string {
	const BaiduHomePageUrl = "https://www.baidu.com"
	return BaiduHomePageUrl
}

func DefaultValueOfString() string {
	var player string
	return player
}

func ValueOfString() string {
	//var player string = "Allen"
	var player = "Allen"
	return player
}

func FormatPlayerName(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}
	var player string
	player = fmt.Sprintf("player name is %s", name)
	return player, nil
}

func ShortDeclare(name string) string {
	player := "player is " + name
	return player
}
