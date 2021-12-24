package interface_pattern

import "fmt"

type Printer interface {
	ToString()
}

type Country3 struct {
	Name string
}

type City3 struct {
	Name string
}

func (c *Country3) ToString() {
	fmt.Printf("Country3 name: %v\n", c.Name)
}

func (c City3) ToString() {
	fmt.Printf("City3 name: %v\n", c.Name)
}

func Poly(p Printer) {
	p.ToString()
}
