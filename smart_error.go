package errors

import (
	"fmt"
	"strings"
)

type SmartError interface {
	error
	ErrorShort() string
	New(parentError error, args ...interface{}) SmartError
	ToOriginal() SmartError
	Traceback() Traceback
	InitialError() SmartError
	ErrorStack() []SmartError
	SetCutOffFirstNLinesOfTraceback(value int) SmartError
}

type smartError struct {
	error
	args      []interface{}
	parent    *smartError
	traceback Traceback
	original  *smartError
}

func argsToStr(args []interface{}) string {
	var argStrs []string
	for _, arg := range args {
		argStrs = append(argStrs, fmt.Sprintf("%v", arg))
	}
	return "[ "+strings.Join(argStrs, " | ")+" ]"
}

func (err smartError) getError() error {
	if err.error != nil {
		return err.error
	}
	if err.parent != nil {
		return err.parent.getError()
	}
	if len(err.args) > 0 {
		initialError, ok := err.args[0].(error)
		if ok {
			return initialError
		}
	}

	return fmt.Errorf("%v", err.args)
}

func (err smartError) ErrorShort() (result string) {
	printError := err.getError()
	if len(err.args) > 0 {
		if err.args[0] == printError {
			err.args = err.args[1:]
		}
	}
	return fmt.Sprintf("%v: %v", err.getError(), argsToStr(err.args))
}

func (err smartError) Error() (result string) {
	errorStack := err.ErrorStack()
	for idx, oneError := range errorStack {
		smartErr := oneError.(*smartError)
		prefix := ""
		if idx > 0 {
			prefix = "caused by: "
		}
		result += prefix + fmt.Sprintf("%v: %v\n", smartErr.getError().Error(), argsToStr(smartErr.args))
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

func (err smartError) SetCutOffFirstNLinesOfTraceback(value int) SmartError {
	err.traceback = err.traceback.(*traceback).setCutOffFirstNLines(value)
	return &err
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

func Wrap(prevErr error, args ...interface{}) SmartError {
	if prevErr == nil {
		return nil
	}

	newErr := *SomeError.(*smartError)
	parentSmartErr, ok := prevErr.(*smartError)
	if ok {
		newErr.parent = parentSmartErr
	} else {
		args = append([]interface{}{prevErr}, args...)
	}

	newErr.args = args
	newErr.traceback = newTraceback()
	newErr.original = SomeError.(*smartError)
	return &newErr
}

func New(prevErr error, args ...interface{}) SmartError {
	newErr := *SomeError.(*smartError)
	parentSmartErr, ok := prevErr.(*smartError)
	if ok {
		newErr.parent = parentSmartErr
	} else {
		args = append([]interface{}{prevErr}, args...)
	}

	newErr.args = args
	newErr.traceback = newTraceback()
	newErr.original = SomeError.(*smartError)
	return &newErr
}
