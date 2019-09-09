package errors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	FormatOneLine = "{%fTC(%fT) %}\"%fE\"{%c>1 caused by: {%lTC(%lT) %}\"%lE\"%}{%clA>0 with args: %lA%}"

	FormatFull = FormatOneLine + "\n" +
		"{%c>1 %aE (%aT)\n%}" +
		"\n    The stack-trace of the initial error:\n" +
		"%lS"
)

var (
	DefaultFormat = FormatFull
)

var (
	formatRegexp_cgt1   = regexp.MustCompile(`{%c>1[^}]*%}`)
	formatRegexp_clAgt0 = regexp.MustCompile(`{%clA>0[^}]*%}`)
	formatRegexp_fTC    = regexp.MustCompile(`{%fTC[^}]*%}`)
	formatRegexp_lTC    = regexp.MustCompile(`{%lTC[^}]*%}`)
)

type Error struct {
	Args         []interface{}
	Err          error
	WrappedError *Error
	Traceback    *Traceback
	Format       string
}

func (err *Error) Is(cmp error) bool {
	return err.Err == cmp
}

func (err *Error) Has(cmp error) bool {
	curErr := err
	for curErr != nil {
		if err.Err == cmp {
			return true
		}
		curErr = curErr.WrappedError
	}
	return false
}

func (err *Error) Deepest() *Error {
	curErr := err
	for curErr.WrappedError != nil {
		curErr = curErr.WrappedError
	}
	return curErr
}

func (err *Error) Unwrap() error {
	if err.WrappedError == nil {
		return err.Err
	}
	return err.WrappedError
}

func reverseStrings(a []string) {
	l := len(a)
	for i := 0; i < l/2; i++ {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
}

func (err *Error) Error() string {
	if err.Traceback == nil {
		return err.Err.Error()
	}

	var replaceOldNew []string

	first := err
	last := err.Deepest()

	for errFmt, errInst := range map[string]*Error{`f`: first, `l`: last} {
		var args string
		if len(errInst.Args) > 0 {
			args = fmt.Sprintf(" %v", errInst.Args)
		}
		var errMsg string
		if smartErr, ok := errInst.Err.(*Error); ok {
			errMsg = smartErr.Err.Error()
		} else {
			errMsg = errInst.Err.Error()
		}
		replaceOldNew = append(replaceOldNew,
			`%`+errFmt+`E`, errMsg,
			`%`+errFmt+`T`, fmt.Sprintf(`%T`, errInst.Err),
			`%`+errFmt+`S`, errInst.Traceback.String(),
			`%`+errFmt+`A`, fmt.Sprint(errInst.Args),
			`% `+errFmt+`A`, args,
		)
	}

	cur := first
	var allError, allType []string
	for cur != nil {
		newErr := cur.Err.Error()
		if len(allError) == 0 || newErr != allError[len(allError)-1] {
			allError = append(allError, newErr)
		}
		newType := fmt.Sprintf(`%T`, cur.Err)
		if len(allType) == 0 || newType != allType[len(allType)-1] {
			allType = append(allType, newType)
		}
		cur = cur.WrappedError
	}
	reverseStrings(allError)
	reverseStrings(allType)
	replaceOldNew = append(replaceOldNew,
		`%aE`, strings.Join(allError, ` -> `),
		`%aT`, strings.Join(allType, ` -> `),
	)

	format := err.Format
	if format == `` {
		format = DefaultFormat
	}
	if fmt.Sprintf(`%T`, first.Err) != `*errors.errorString` {
		format = formatRegexp_fTC.ReplaceAllString(format, `${1}`)
	} else {
		format = formatRegexp_fTC.ReplaceAllString(format, ``)
	}
	if fmt.Sprintf(`%T`, last.Err) != `*errors.errorString` {
		format = formatRegexp_lTC.ReplaceAllString(format, `${1}`)
	} else {
		format = formatRegexp_lTC.ReplaceAllString(format, ``)
	}
	if first != last {
		format = formatRegexp_cgt1.ReplaceAllString(format, `${1}`)
	} else {
		format = formatRegexp_cgt1.ReplaceAllString(format, ``)
	}
	if len(last.Args) > 0 {
		format = formatRegexp_clAgt0.ReplaceAllString(format, `${1}`)
	} else {
		format = formatRegexp_clAgt0.ReplaceAllString(format, ``)
	}
	return strings.NewReplacer(replaceOldNew...).Replace(format)
}

func (err *Error) Wrap(args ...interface{}) (result *Error) {
	var argErr error
	if len(args) > 0 {
		argErr, _ = args[0].(error)
	}
	if argErr == nil {
		argErr = errors.New(fmt.Sprint(args...))
	} else {
		args = args[1:]
	}

	result = &[]Error{*err}[0]

	result.Args = args
	result.Traceback = newTraceback()

	if wrappedError, ok := argErr.(*Error); ok && wrappedError.Err != nil {
		result.Err = err
		result.WrappedError = wrappedError
		return
	}
	if result.Err == nil {
		result.Err = argErr
		return
	}

	result.Err = err
	result.WrappedError = &Error{
		Err:       argErr,
		Traceback: result.Traceback,
	}
	return
}

func Wrap(prevErr error, args ...interface{}) *Error {
	if prevErr == nil {
		return nil
	}

	var argErr error
	if len(args) > 0 {
		argErr, _ = args[0].(error)

		if argErr == nil {
			argErr = errors.New(fmt.Sprint(args...))
		} else {
			args = args[1:]
		}
	}

	err, _ := argErr.(*Error)
	if err == nil {
		err = &[]Error{*UndefinedError}[0]
	}

	return err.Wrap(append([]interface{}{prevErr}, args...)...)
}

func New(err interface{}, args ...interface{}) *Error {
	newErr := &[]Error{*UndefinedError}[0]
	newErr.Err, _ = err.(error)
	if newErr.Err == nil {
		newErr.Err = fmt.Errorf("%v", err)
	}
	newErr.Args = args
	newErr.Traceback = newTraceback()

	return newErr
}
