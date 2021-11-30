package main

import "fmt"

func main() {
	//exampleNil1()
	//exampleNil2()
	//exampleNil3()
	//exampleNil4()
	//exampleNil5()
	//exampleNil6()
	exampleNil7()
}

func exampleNil1() {
	var configMap map[string]interface{}
	configMap["name"] = "luca"

	// Output
	// panic: assignment to entry in nil map
}

func exampleNil2() {
	var configList []string
	configList[0] = "1"

	// Output:
	// panic: runtime error: index out of range [0] with length 0
}

func exampleNil3() {
	configList := make([]string, 1)
	configList[0] = "1"
	fmt.Println(configList)

	// Output:
	// [1]
}

// append函数会生成新的切片, 在底层为切片分配了数组
func exampleNil4() {
	var configList []string
	configList = append(configList, "1")
	fmt.Println(configList)

	// Output:
	// [1]
}

func exampleNil5() {
	var configList1 []string // TODO 没数据的时候是Null，这也是切片未初始化导致的 ==> json
	configList2 := make([]string, 1)
	fmt.Printf("configList1 ==> %v\n", configList1)
	fmt.Printf("configList2 ==> %v\n", configList2)

	// Output:
	// configList1 ==> []
	// configList2 ==> []
}

type Oops struct {
	value string
}

func (o *Oops) toString() string {
	return o.value
}

// 指针的零值nil与*T并不相等 所以指针类型的变量在使用前要注意使用new函数进行初始化 ==> 查看exampleNil7
func exampleNil6() {
	var oops *Oops
	_ = oops.value
	// or
	//oops.toString()

	// Output:
	// panic: runtime error: invalid memory address or nil pointer dereference
}

func exampleNil7() {
	var oops = new(Oops)
	oops.toString()

	// or
	var oops1 Oops
	oops1.toString()
}
