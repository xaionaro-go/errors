package errors

import (
	"fmt"
)

var (
	UndefinedError        = &Error{Err: nil}
	InvalidArguments      = &Error{Err: fmt.Errorf(`invalid arguments`)}
	NotImplemented        = &Error{Err: fmt.Errorf(`not implemented (yet?)`)}
	ProtocolMismatch      = &Error{Err: fmt.Errorf(`protocol mismatch`)}
	NotFound              = &Error{Err: fmt.Errorf(`not found`)}
	OutOfRange            = &Error{Err: fmt.Errorf(`out of range`)}
	CannotResolveAddress  = &Error{Err: fmt.Errorf(`cannot resolve the address`)}
	CannotWriteToFile     = &Error{Err: fmt.Errorf(`cannot write to the file`)}
	CannotParseFile       = &Error{Err: fmt.Errorf(`cannot parse the file`)}
	CannotOpenFile        = &Error{Err: fmt.Errorf(`cannot open the file`)}
	CannotSendData        = &Error{Err: fmt.Errorf(`cannot send the data`)}
	CannotSetRLimitNoFile = &Error{Err: fmt.Errorf(`cannot set limit "nofile"`)}
	UnableToConnect       = &Error{Err: fmt.Errorf(`unable to connect`)}
	UnableToGetKey        = &Error{Err: fmt.Errorf(`unable to get a key`)}
	UnableToStartSession  = &Error{Err: fmt.Errorf(`unable to start a session`)}
	UnableToListen        = &Error{Err: fmt.Errorf(`unable to start listening`)}
	UnableToParse         = &Error{Err: fmt.Errorf(`unable to parse`)}
	UnableToFetchData     = &Error{Err: fmt.Errorf(`unable to fetch the data`)}
	UnableToProcessData   = &Error{Err: fmt.Errorf(`unable to process the data`)}
	UnexpectedInput       = &Error{Err: fmt.Errorf(`unexpected input`)}
	DBNotInitialized      = &Error{Err: fmt.Errorf(`DB is not initialized`)}
)
