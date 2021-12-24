package main

import "fmt"

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("2")
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

type OtherThingsVisitor struct {
	visitor Visitor
}

func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("3")
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("1")
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	//info := Info{}
	//var v Visitor = &info
	//v = LogVisitor{v}
	//v = NameVisitor{v}
	//v = OtherThingsVisitor{v}
	//loadFile := func(info *Info, err error) error {
	//	fmt.Println("4")
	//	info.Name = "Hao Chen"
	//	info.Namespace = "MegaEase"
	//	info.OtherThings = "We are running as remote team."
	//	return nil
	//}
	//_ = v.Visit(loadFile)

	info := Info{}
	var v Visitor = &info

	loadFile := func(info *Info, err error) error {
		fmt.Println("4")
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v = NewDecoratedVisitor(v, LogVisitorFunc, NameVisitorFunc, OtherVisitorFunc)
	_ = v.Visit(loadFile)
}

func LogVisitorFunc(info *Info, err error) error {
	fmt.Println("1")
	fmt.Println("LogVisitor() before call function")
	fmt.Println("LogVisitor() after call function")
	return err
}

func NameVisitorFunc(info *Info, err error) error {
	fmt.Println("2")
	fmt.Println("NameVisitor() before call function")
	if err == nil {
		fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	}
	fmt.Println("NameVisitor() after call function")
	return err
}

func OtherVisitorFunc(info *Info, err error) error {
	fmt.Println("3")
	fmt.Println("OtherThingsVisitor() before call function")
	if err == nil {
		fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	}
	fmt.Println("OtherThingsVisitor() after call function")
	return err
}
