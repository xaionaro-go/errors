package errors

import (
	"fmt"
)

var (
	NotImplemented        SmartError = &smartError{error: fmt.Errorf(`Not implemented (yet?)`)}
	UnableToConnect       SmartError = &smartError{error: fmt.Errorf(`Unable to connect`)}
	ProtocolMismatch      SmartError = &smartError{error: fmt.Errorf(`Protocol mismatch`)}
	NotFound              SmartError = &smartError{error: fmt.Errorf(`Not found`)}
	OutOfRange            SmartError = &smartError{error: fmt.Errorf(`Out of range`)}
	CannotResolveAddress  SmartError = &smartError{error: fmt.Errorf(`Cannot resolve the address`)}
	CannotWriteToFile     SmartError = &smartError{error: fmt.Errorf(`Cannot write to the file`)}
	CannotParseFile       SmartError = &smartError{error: fmt.Errorf(`Cannot parse the file`)}
	CannotOpenFile        SmartError = &smartError{error: fmt.Errorf(`Cannot open the file`)}
	CannotSendData        SmartError = &smartError{error: fmt.Errorf(`Cannot send the data`)}
	CannotSetRLimitNoFile SmartError = &smartError{error: fmt.Errorf(`Cannot set limit "nofile"`)}
	CannotFetchData       SmartError = &smartError{error: fmt.Errorf(`Cannot fetch the data`)}
	UnableToGetKey        SmartError = &smartError{error: fmt.Errorf(`Unable to get a key`)}
	UnableToStartSession  SmartError = &smartError{error: fmt.Errorf(`Unable to start a session`)}
	UnableToListen        SmartError = &smartError{error: fmt.Errorf(`Unable to start listening`)}
	UnableToParse         SmartError = &smartError{error: fmt.Errorf(`Unable to parse`)}
	UnableToFetchData     SmartError = &smartError{error: fmt.Errorf(`Unable to fetch the data`)}
	UnexpectedInput       SmartError = &smartError{error: fmt.Errorf(`Unexpected input`)}
)
