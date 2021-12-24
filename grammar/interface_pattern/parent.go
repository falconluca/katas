package interface_pattern

import "fmt"

type WithName struct {
	Name string
}

type Country struct {
	WithName // 结构体内嵌、委托、Shared struct
}

type City struct {
	WithName
}

type Printable interface {
	PrintStr()
}

func (p *WithName) PrintStr() {
	fmt.Printf("name: %v\n", p.Name)
}
