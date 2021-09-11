package ioc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmbedStruct(t *testing.T) {
	assert2 := assert.New(t)

	boxModel := BoxModel{Height: 10, Width: 401}
	navbar := Label{Text: "网站导航栏", BoxModel: boxModel}

	assert2.Equal(10, navbar.Height) // 哈哈不知道为什么，觉得有点像面向对象的继承
	assert2.Equal(10, navbar.BoxModel.Height)

	assert2.Equal(401, navbar.Width)
	assert2.Equal(401, navbar.BoxModel.Width)

	assert2.Equal("网站导航栏", navbar.Text)
}
