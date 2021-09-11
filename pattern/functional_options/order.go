package functional_options

type OrderStatus struct {
	code int
	msg  string
}

type Order struct {
	OrderId string
	GoodsId string
	Status  OrderStatus
}

// 函数类型
type OrderOption func(*Order)

// 三个高阶函数
//func OrderId(orderId string) OrderOption {
//	return func(order *Order) {
//		order.OrderId = orderId
//	}
//}

func GoodsId(goodsId string) OrderOption {
	return func(order *Order) {
		order.GoodsId = goodsId
	}
}

func Status(status OrderStatus) OrderOption {
	return func(order *Order) {
		order.Status = status
	}
}

// 创建
func NewOrder(orderId string, options ...func(*Order)) (*Order, error) {
	order := Order{
		// required
		OrderId: orderId,
	}
	for _, option := range options {
		option(&order)
	}
	return &order, nil
}
