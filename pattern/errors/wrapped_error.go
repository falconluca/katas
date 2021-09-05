package errors

import (
	"fmt"
	// https://github.com/pkg/errors
	"github.com/pkg/errors"
)

type AuthorizationError struct {
	operation string
	err       error // original error
}

func (e *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

func WrappedError() error {
	// 错误包装
	var err error
	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	// Cause接口
	switch err := errors.Cause(err).(type) {
	case *AuthorizationError:
		// handle specifically
		return &AuthorizationError{err: err, operation: "401"}
	default:
		// unknown error
		return nil
	}
}
