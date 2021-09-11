package functional_options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMacBook(t *testing.T) {
	assert := assert.New(t)

	cpu := &Cpu{Core: 8, CType: "M1"}
	macBook, _ := NewDefaultMacBook("Luca's Mac", 13.3, cpu)
	assert.Equal(float32(13.3), macBook.Size)
	assert.Equal("M1", macBook.Cpu.CType)
}

func TestNewDefaultMacBook(t *testing.T) {
	assert := assert.New(t)

	macBook, _ := NewDefaultMacBook("Luca's Mac", 13.3, nil)
	assert.Equal("M2", macBook.Cpu.CType)
}
