package main

import (
	"fmt"
	"testing"

	_ "init/p1"
)

// init/p1下的p1.go的init方法【会】被执行
func TestInit1(t *testing.T) {
	fmt.Println("exec TestInit1")
}
