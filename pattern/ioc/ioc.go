package ioc

import (
	"errors"
	"fmt"
)

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}

	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

// ------------------------

type IntSetIoc struct {
	data map[int]bool
	undo Undo
}

func NewIntSetIoc() IntSetIoc {
	return IntSetIoc{data: make(map[int]bool)}
}

func (set *IntSetIoc) Undo() error {
	return set.undo.Undo()
}

func (set *IntSetIoc) Contains(x int) bool {
	return set.data[x]
}

func (set *IntSetIoc) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntSetIoc) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}

func AfterIoc() {
	intSet := NewIntSetIoc()
	//intSet.Add(78)
	//intSet.Add(74)
	if err := intSet.Undo(); err != nil {
		_ = fmt.Errorf("err: %v", err)
	}
	intSet.Add(77)
	fmt.Println(intSet.data)
}
