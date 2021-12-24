package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

func main() {
	var r *strings.Reader = strings.NewReader("Greetings!")
	b, _ := ioutil.ReadAll(r)
	fmt.Printf("%s\n", b) // Auto convert bytes to string.

	//var f multipart.File // type File interface { io.Reader
	//_, _ = ioutil.ReadAll(f)

	fmt.Println(strconv.Itoa(12))

	// 获取文件后缀
	fmt.Println(path.Ext("/usr/local/webpack.conf.js"))

	// TODO ioutil 查看Go标准库里的Example代码和_test代码
}
