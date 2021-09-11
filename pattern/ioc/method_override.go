package ioc

import "fmt"

/*

委托关系：

	BoxModel
	  /  \
  Label   List
	|
  Button

*/

type Button struct {
	Label
}

type List struct {
	BoxModel
	Texts []string
	Index int
}

// == interface ==

type Painter interface {
	Paint() string
}

type Clickable interface {
	Click() string
}

// == label ==

func (label Label) Paint() string {
	return fmt.Sprintf("Label paint %s", label.Text)
}

// == button ==

func (button Button) Click() string {
	return fmt.Sprintf("Button click %s", button.Text)
}

func (button Button) Paint() string {
	return fmt.Sprintf("Button paint %s", button.Text)
}

// == List ==

func (list List) Click() string {
	return fmt.Sprintf("List click %s", list.Texts[0])
}

func (list List) Paint() string {
	return fmt.Sprintf("List paint %s", list.Texts[0])
}
