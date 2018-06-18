package errors

import (
	"runtime/debug"
	"strings"
)

type Traceback interface {
	String() string
}

type traceback struct {
	data []byte
}

func newTraceback() *traceback {
	return &traceback{data: debug.Stack()}
}

func (traceback traceback) String() string {
	stackString := string(traceback.data)
	stackLines := strings.Split(stackString, "\n")
	return strings.Join(stackLines[7:], "\n")
}
