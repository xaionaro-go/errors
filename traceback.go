package errors

import (
	"runtime/debug"
	"strings"
)

const (
	cutOffFirstNLinesOfTraceback = 7
)

type Traceback interface {
	String() string
}

type traceback struct {
	cutOffFirstNLines int
	data []byte
}

func newTraceback() *traceback {
	return &traceback{data: debug.Stack(), cutOffFirstNLines: cutOffFirstNLinesOfTraceback}
}

func (traceback traceback) setCutOffFirstNLines(cutOffFirstNLines int) *traceback {
	traceback.cutOffFirstNLines = cutOffFirstNLines
	return &traceback
}

func (traceback traceback) String() string {
	stackString := string(traceback.data)
	stackLines := strings.Split(stackString, "\n")
	return strings.Join(stackLines[traceback.cutOffFirstNLines:], "\n")
}
