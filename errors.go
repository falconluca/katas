package main

import (
	"errors"
	"fmt"
)

// 自定义异常 type error interface { Error() string }

type ApiError struct {
	statusCode int
	errMsg     string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("msg: %s code: %v", e.errMsg, e.statusCode)
}

func GenerateHttpError(statusCode int) (int, error) {
	if statusCode >= 200 && statusCode < 300 {
		return 0, errors.New("OK")
	} else if statusCode >= 400 && statusCode < 500 {
		return 0, &ApiError{statusCode, "Bad Request"}
	} else {
		return statusCode * statusCode, nil
	}
}
