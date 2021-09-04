package main

func SayHey(name string) string {
	return "Hey " + name
}

func GetOrderStatusById(orderId, userId int) (string, bool) {
	orderStatusTable := map[int]string{0: "unpaid", 1: "paid", 2: "finished"}
	orderStatus, err := orderStatusTable[orderId]
	return orderStatus, err
}

func SwapString(x string, y string) (before, after string) {
	before = y
	after = x
	return
}

func CalculateTotalGoodsStock(goodsStockList ...int) int {
	var result int
	// 使用_符号忽略列表下标
	for _, goodsStock := range goodsStockList {
		//fmt.Printf("idx: %v\n", idx)
		result += goodsStock
	}
	return result
}

// Fb 斐波那契数列
func Fb(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fb(n-1)
}
