package errors

import (
	"fmt"
)

type SmartError interface {
	error
	New(parentError error, args ...interface{}) SmartError
	ToOriginal() SmartError
	Traceback() Traceback
	InitialError() SmartError
	ErrorStack() []SmartError
}

type smartError struct {
	error
	args      []interface{}
	parent    *smartError
	traceback Traceback
	original  *smartError
}

func (err smartError) Error() (result string) {
	errorStack := err.ErrorStack()
	for idx, oneError := range errorStack {
		smartErr := oneError.(*smartError)
		prefix := ""
		if idx > 0 {
			prefix = "caused by: "
		}
		result += prefix + fmt.Sprintf("%v: %v\n", smartErr.error.Error(), smartErr.args)
	}
	traceback := errorStack[len(errorStack)-1].Traceback()
	if traceback != nil {
		result += "The traceback of the initial error:\n" + traceback.String()
	}
	return
}

func (err *smartError) New(prevErr error, args ...interface{}) SmartError {
	newErr := *err
	parentSmartErr, ok := prevErr.(*smartError)
	if ok {
		newErr.parent = parentSmartErr
	} else {
		args = append([]interface{}{prevErr}, args...)
	}

	newErr.args = args
	newErr.traceback = newTraceback()
	newErr.original = err
	return &newErr
}

func (err smartError) ErrorStack() (result []SmartError) {
	var errPointer, parentError *smartError
	errPointer = &err
	for parentError = errPointer.parent; parentError != nil; {
		result = append(result, errPointer)
		errPointer = parentError
		parentError = errPointer.parent
	}
	result = append(result, errPointer)
	return
}

func (err smartError) InitialError() SmartError {
	errorStack := err.ErrorStack()
	return errorStack[len(errorStack)-1]
}

func (err smartError) Traceback() Traceback {
	initialErr := err.InitialError().(*smartError)
	if initialErr == nil {
		return nil
	}
	return initialErr.traceback
}

func (err smartError) ToOriginal() SmartError {
	return err.original
}
