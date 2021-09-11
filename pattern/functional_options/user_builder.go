package functional_options

import (
	"errors"
	"strings"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Gender     string // fixme
	IsActivate bool
	// TODO 在处理错误的时候可能就有点麻烦（需要为Server结构增加一个error 成员，破坏了Server结构体的“纯洁”）
}

type UserBuilder struct {
	User
}

func (userBuilder *UserBuilder) Builder() *UserBuilder {
	// required values
	return userBuilder
}

func (userBuilder *UserBuilder) Username(username string) *UserBuilder {
	userBuilder.User.Username = username
	return userBuilder
}

func (userBuilder *UserBuilder) Gender(gender string) *UserBuilder {
	userBuilder.User.Gender = gender
	return userBuilder
}

func (userBuilder *UserBuilder) Password(password string) *UserBuilder {
	userBuilder.User.Password = password
	return userBuilder
}

func (userBuilder *UserBuilder) Id(id int) *UserBuilder {
	userBuilder.User.Id = id
	return userBuilder
}

func (userBuilder *UserBuilder) IsActivate(isActivate bool) *UserBuilder {
	userBuilder.User.IsActivate = isActivate
	return userBuilder
}

func (userBuilder *UserBuilder) Build() (User, error) {
	// validate params... fixme do one thing at a time.
	gender := userBuilder.User.Gender
	if strings.Trim(gender, " ") == "" {
		return userBuilder.User, errors.New("gender cannot be blank")
	}
	if strings.EqualFold(gender, "Male") || strings.EqualFold(gender, "Female") {
		return userBuilder.User, nil
	}
	return userBuilder.User, errors.New("gender error")
}
