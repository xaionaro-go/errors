See also [github.com/pkg/errors](https://github.com/pkg/errors).

`example/main.go`:
```go
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
```

```
$ go run example/main.go
panic: Cannot send the data: [ hello ]
caused by: Cannot write to the file: [ hello ]
caused by: Cannot open the file: [ open /wrong/path/hello.txt: no such file or directory | some comment here | also we can pass the file path, for example | /wrong/path/hello.txt ]
The traceback of the initial error:
main.writeToLogFile(0x4b8faa, 0x5, 0x0, 0x0)
        /home/xaionaro/gocode/src/github.com/xaionaro-go/errors/example/main.go:18 +0x139
main.sendString(0x4b8faa, 0x5, 0xc420012401, 0xc42000e2a0)
        /home/xaionaro/gocode/src/github.com/xaionaro-go/errors/example/main.go:29 +0x4d
main.doTheMagic(0x19, 0x19)
        /home/xaionaro/gocode/src/github.com/xaionaro-go/errors/example/main.go:38 +0x3a
main.main()
        /home/xaionaro/gocode/src/github.com/xaionaro-go/errors/example/main.go:46 +0x26


goroutine 1 [running]:
main.main()
        /home/xaionaro/gocode/src/github.com/xaionaro-go/errors/example/main.go:49 +0x19c
exit status 2
```
