package main

import (
	"fmt"
	"init/p1"
	"testing"
)

// init/p1下的p1.go的init方法【会】被执行
func TestInit3(t *testing.T) {
	fmt.Println("exec TestInit2")
	p1.Run()
}

// TODO 多个init方法的执行顺序
