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
	NotImplemented = smartError{error: fmt.Errorf("Not implemented (yet?)")}
	UnableToConnect = smartError{error: fmt.Errorf("Unable to connect")}
	ProtocolMismatch = smartError{error: fmt.Errorf("Protocol mismatch")}
	NotFound = smartError{error: fmt.Errorf("Not found")}
	OutOfRange = smartError{error: fmt.Errorf("Out of range")}
	CannotResolveAddress = smartError{error: fmt.Errorf("Cannot resolve the address")}
	CannotWriteToFile = smartError{error: fmt.Errorf("Cannot write to the file")}
	CannotParseFile = smartError{error: fmt.Errorf("Cannot parse the file")}
	CannotOpenFile = smartError{error: fmt.Errorf("Cannot open the file")}
	UnableToGetKey = smartError{error: fmt.Errorf("Unable to get a key")}
	UnableToStartSession = smartError{error: fmt.Errorf("Unable to start a session")}
	UnableToListen = smartError{error: fmt.Errorf("Unable to start listening")}
)
