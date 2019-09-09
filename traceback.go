package errors

import (
	"runtime/debug"
	"strings"
)

const (
	cutOffFirstNLinesOfTraceback = 7
)

type Traceback struct {
	CutOffFirstNLines int
	Data              []byte
}

func newTraceback() *Traceback {
	return &Traceback{Data: debug.Stack(), CutOffFirstNLines: cutOffFirstNLinesOfTraceback}
}

func (traceback Traceback) String() string {
	stackString := string(traceback.Data)
	stackLines := strings.Split(stackString, "\n")
	return strings.Join(stackLines[traceback.CutOffFirstNLines:], "\n")
}
