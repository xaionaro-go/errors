package errors

import (
	"fmt"
	"strings"
)

func Wrap(prevErr error, args ...interface{}) Interface {
	if prevErr == nil {
		return nil
	}

	var err *Error
	if len(args) > 0 {
		err, _ = args[0].(*Error)
		if err != nil {
			if err.Traceback != nil {
				err = nil
			}
		}
		if err != nil {
			args = args[1:]
		}
	}

	result := err.Wrap(append([]interface{}{prevErr}, args...)...).(*Error)
	result.Traceback.CutOffFirstNLines += 2
	if prevErrCasted, ok := prevErr.(*Error); ok {
		result.Format = prevErrCasted.Format
	}
	return result
}

// Errorf is analog of fmt.Errorf to wrap errors, but
// it creates this smart errors with tracebacks.
func Errorf(format string, args ...interface{}) error {
	errI := args[len(args)-1]
	err, _ := errI.(error)
	if errI != nil && err == nil {
		return nil // nil error, nothing to wrap
	}

	if err == nil {
		return New(fmt.Errorf(format, args...))
	}
	args = args[:len(args)-1]

	var parentErr error
	if strings.HasSuffix(format, ": %w") {
		parentErr = fmt.Errorf(format[:len(format)-4], args...)
	} else {
		parentErr = fmt.Errorf(format, args...)
	}
	parentErrForWrap := New(parentErr)
	parentErrForWrap.Traceback = nil
	return Wrap(err, parentErrForWrap)
}
