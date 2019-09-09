package main

import (
	"fmt"
	"os"

	"github.com/xaionaro-go/errors"
)

const (
	HELLO_STRING = "hello"
	LOGFILE_PATH = "/wrong/path/hello.txt"
)

var (
	ErrCannotDoMagic     = errors.New(`cannot do magic`)
	ErrCannotDoDeepMagic = errors.New(`cannot do deep magic`)
)

func writeToLogFile(str string) (err error) {
	defer func() { err = errors.Wrap(err, LOGFILE_PATH, str) }()

	f, err := os.OpenFile(LOGFILE_PATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		return
	}
	return nil
}

func sendString(str string) error {
	err := writeToLogFile(str)
	if err != nil {
		return errors.New(`unable to send string`).Wrap(err, str)
	}
	return nil
}

func doEvenMoreMagic() error {
	err := sendString(HELLO_STRING)
	if err != nil {
		return errors.CannotSendData.Wrap(err, HELLO_STRING)
	}
	return nil
}
func doMoreMagic() error {
	err := doEvenMoreMagic()
	if err != nil {
		return errors.Wrap(err, ErrCannotDoDeepMagic, `more comments`)
	}
	return nil
}

func doTheMagic() error {
	err := doMoreMagic()
	if err != nil {
		return ErrCannotDoMagic.Wrap(err, `some comments`)
	}
	return nil
}

func main() {
	err := doTheMagic()
	switch err.(*errors.Error).Err {
	case ErrCannotDoMagic:
		panic(err)
	case errors.ProtocolMismatch:
		fmt.Println("We will never get here")
	}

	fmt.Println("This shouldn't happened")
}
