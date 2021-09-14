package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmp(t *testing.T) {
	assertions := assert.New(t)

	var employeeList = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
		{"Marry", 29, 0, 6000},
		{"Mike", 32, 8, 4000},
	}
	// 1.统计有多少员工大于40岁
	oldCount := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Age > 40
	})
	assertions.Equal(2, oldCount)

	// 2.统计有多少员工薪水大于6000
	highPay := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	assertions.Equal(4, highPay)

	// 3.列出有没有休假的员工
	noVacation := EmployeeFilterIn(employeeList, func(e *Employee) bool {
		return e.Vacation == 0
	})
	assertions.Equal(3, len(noVacation))

	// 4.统计所有员工的薪资总和
	totalPay := EmployeeSumIf(employeeList, func(e *Employee) int {
		return e.Salary
	})
	assertions.Equal(43500, totalPay)

	// 5.统计30岁以下员工的薪资总和
	youngerPay := EmployeeSumIf(employeeList, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	assertions.Equal(19000, youngerPay)
}
