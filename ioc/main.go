package main

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	data := make(map[int]bool)
	return IntSet{data: data}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

func main() {
	//intSet := NewIntSet()
	//intSet.Add(12)

	type Test struct{}
	v := Test{}
	Print(v)
}

func Print(v interface{}) {
	println(v)
}
