package map_reduce

type Employee struct {
	Name string
	Age  int
	// 休假
	Vacation int
	Salary   int
}

func EmployeeCountIf(
	// 数据
	list []Employee,
	// 业务操作
	fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) { // 计数
			count += 1
		}
	}
	return count
}

func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newList []Employee
	for i, _ := range list {
		if fn(&list[i]) { // 过滤
			newList = append(newList, list[i])
		}
	}
	return newList
}

func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list { // 计算和
		sum += fn(&list[i])
	}
	return sum
}
