package gin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetMapReturnValues(t *testing.T) {
	assertions := assert.New(t)

	table := make(map[string]string)
	table["name"] = "luca"

	_, ok := table["name"]
	assertions.True(ok)
	_, ok = table["age"]
	assertions.False(ok)
}

func TestGetCurrentPath(t *testing.T) {
	// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
	dir, _ := os.Getwd()
	fmt.Println(dir) // /Users/luca/dev-resources/katas/gin
}
