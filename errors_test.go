package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	result, e := GenerateHttpError(10)
	assert.Equal(t, 100, result)
	assert.Empty(t, e)

	result, e = GenerateHttpError(201)
	assert.Equal(t, 0, result)
	assert.NotEmpty(t, e)
	assert.Equal(t, "OK", e.Error())

	result, e = GenerateHttpError(401)
	assert.Equal(t, 0, result)
	assert.NotEmpty(t, e)
	assert.Equal(t, "msg: Bad Request code: 401", e.Error())
}
