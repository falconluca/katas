package ioc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMethodOverride(t *testing.T) {
	assertion := assert.New(t)

	label := Label{"Login", BoxModel{10, 70}}
	assertion.Equal("Label paint Login", label.Paint())

	loginButton := Button{Label{"Login", BoxModel{10, 70}}}
	assertion.Equal("Button paint Login", loginButton.Paint()) // 复写了Label的Paint
	assertion.Equal("Button click Login", loginButton.Click())

	list := List{BoxModel{10, 70}, []string{"数据源1"}, 0}
	assertion.Equal("List paint 数据源1", list.Paint())
	assertion.Equal("List click 数据源1", list.Click())
}
