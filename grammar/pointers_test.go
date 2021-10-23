package grammar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointers(t *testing.T) {
	var devSalary int = 10000
	// 初始化指针(获取变量的指针)
	var pLucaSalary *int = &devSalary // 举例: 0xc0000a2b00
	assert.Equal(t, devSalary, *pLucaSalary)

	// 访问指针(取值get或赋值set)
	*pLucaSalary = 12000
	assert.Equal(t, devSalary, *pLucaSalary)

	devSalary = 15000
	assert.Equal(t, devSalary, *pLucaSalary)
}
