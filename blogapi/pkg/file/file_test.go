package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheckPermission(t *testing.T) {
	a := assert.New(t)

	wd, _ := os.Getwd()
	ok := HasPermission(wd)
	a.True(ok)
}
