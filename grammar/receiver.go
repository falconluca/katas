package grammar

import "fmt"

// TOOD 两种结构体内嵌的方式

type House struct {
	HouseId int
	Address string
}

// ---------- Function -------------

func printHouse(h *House) {
	fmt.Printf("origin house: %+v\n", *h)
	h.Address = fmt.Sprintf("%s(printed)", h.Address)
	fmt.Printf("printed house: %+v\n", *h)
}

func Function() {
	house := &House{HouseId: 10001, Address: "万达"}
	printHouse(house)
}

// ---------- Receiver -------------

func (h *House) Print() {
	fmt.Printf("origin house: %+v\n", *h)
	h.Address = fmt.Sprintf("%s(printed)", h.Address)
	fmt.Printf("printed house: %+v\n", *h)
}

func Receiver() {
	house := &House{HouseId: 10002, Address: "建发"}
	house.Print()
}
