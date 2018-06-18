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

func writeToLogFile(str string) error {
	f, err := os.OpenFile(LOGFILE_PATH, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
	if err != nil {
		return errors.CannotOpenFile.New(err, "some comment here", "also we can pass the file path, for example", LOGFILE_PATH)
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		return errors.CannotWriteToFile.New(err, str)
	}
	return nil
}

func sendString(str string) error {
	err := writeToLogFile(str)
	if err != nil {
		return errors.CannotWriteToFile.New(err, str)
	}
	fmt.Println(str)
	return nil
}

func doTheMagic() error {
	err := sendString(HELLO_STRING)
	if err != nil {
		return errors.CannotSendData.New(err, HELLO_STRING)
	}
	return nil
}

func main() {
	err := doTheMagic()
	switch err.(errors.SmartError).ToOriginal() {
	case errors.CannotSendData:
		panic(err)
	case errors.ProtocolMismatch:
		fmt.Println("We will never get here")
	}

	fmt.Println("This shouldn't happened")
}
