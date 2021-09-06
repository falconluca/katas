package ioc

import (
	"errors"
	"fmt"
)

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	data := make(map[int]bool)
	return IntSet{data}
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

// Undo ...

type UndoableIntSet struct { // Poor style
	IntSet    // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) { // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
		return
	}

	set.functions = append(set.functions, nil)
}

func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
		return
	}

	set.functions = append(set.functions, nil)
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("no functions to undo")
	}

	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

// -------ioc

func BeforeIoc() {
	newIntSet := NewIntSet()
	newIntSet.Add(0)
	newIntSet.Add(2)
	newIntSet.Add(7)
	fmt.Println(newIntSet)
	newIntSet.Delete(2)
	fmt.Println(newIntSet)
	fmt.Println(newIntSet.Contains(7))
}
