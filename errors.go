package errors

import (
	"fmt"
)

type smartError struct {
	error
	args []interface{}
}

func (err smartError) String() string {
	return fmt.Sprintf("%v: %v", err.error.Error(), err.args)
}

func (err smartError) SetArgs(args ...interface{}) *smartError {
	err.args = args
	return &err
}

func (err smartError) ToError() error {
	return err.error
}

var (
	ErrNotImplemented = smartError{error: fmt.Errorf("Not implemented (yet?)")}
)
