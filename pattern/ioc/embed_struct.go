package ioc

type BoxModel struct {
	Height int
	Width  int
}

type Label struct {
	Text     string // Aggregation
	BoxModel        // 内嵌结构体 -> 实际上就是委托(delegation)
}
