package main

import (
	"fmt"
	"testing"
)

// init/p1下的p1.go的init方法【不会】被执行
func TestInit2(t *testing.T) {
	fmt.Println("exec TestInit2")
}
