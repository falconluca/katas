package grammar

import (
	"errors"
	"fmt"
)

// Go的基本类型 https://tour.go-zh.org/basics/11

func GetConst() (string, bool) {
	const BaiduHomePageUrl = "https://www.baidu.com"
	// 常量不能用 短变量 语法声明
	const Enable = true
	// TODO 数值常量, 是高精度的值
	return BaiduHomePageUrl, Enable
}

func DefaultValues() (string, int, bool) {
	var player string
	var number int
	var isStarting bool
	return player, number, isStarting
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
