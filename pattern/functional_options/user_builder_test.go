package functional_options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyGenderUserBuilder(t *testing.T) {
	assert := assert.New(t)

	userBuilder := UserBuilder{}
	_, err := userBuilder.Builder().Build()
	assert.Equal("gender cannot be blank", err.Error())
	//assert.Equal("gender error", err)
}

func TestErrorGenderUserBuilder(t *testing.T) {
	assert := assert.New(t)

	userBuilder := UserBuilder{}
	_, err := userBuilder.Builder().
		Gender("Ma;e").
		Build()
	assert.Equal("gender error", err.Error())
}

func TestUserBuilder(t *testing.T) {
	assert := assert.New(t)

	userBuilder := UserBuilder{}
	user, err := userBuilder.Builder().
		Gender("Male").
		Username("Luca").
		Password("123").
		Id(111111).
		Build()
	assert.Nil(err)
	assert.Equal("Male", user.Gender)
	assert.Equal("Luca", user.Username)
	assert.Equal(false, user.IsActivate)
}
