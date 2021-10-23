package grammar

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(test *testing.T) {
	t := time.Now()
	assert.NotEmpty(test, t)

	//assert.Equal(test, "Saturday", t.Weekday().String())
	//assert.Equal(test, time.Saturday, t.Weekday())

	assert.Equal(test, "Saturday", time.Saturday.String())
	assert.Equal(test, "Sunday", time.Sunday.String())
}
