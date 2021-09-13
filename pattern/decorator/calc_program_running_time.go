package decorator

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// SumFunc 抽取了Sum1和Sum2的函数签名
type SumFunc func(int64, int64) int64

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

// getFunctionName 通过反射获取函数名称
func getFunctionName(i interface{}) string {
	pointerOfInterface := reflect.ValueOf(i).Pointer()
	funcF := runtime.FuncForPC(pointerOfInterface)
	return funcF.Name()
}

// timedSumFunc 装饰器
func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 { // 高阶函数
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",
				getFunctionName(f), // 获取被装饰的函数(f)的名称
				time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

// CalcProgramRunningTime 通过函数柯里化"药水"赋予函数计算执行时间的能力
func CalcProgramRunningTime() {
	xSum1 := timedSumFunc(Sum1)
	xSum2 := timedSumFunc(Sum2)

	sum1Result := xSum1(-10000, 10000000)
	fmt.Printf("被增强后的Sum1:  %d\n", sum1Result)

	sum2Result := xSum2(-10000, 10000000)
	fmt.Printf("被增强后的Sum2: %d\n", sum2Result)
}
