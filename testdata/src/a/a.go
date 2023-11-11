package a

import (
	"errors"
	"log"
)

var ErrSomeWrong = errors.New("somethis is wrong")
var SomeWrongError = errors.New("somethis is wrong") // want `Error variable SomeWrongError does not follow naming conventions, rule is \^Err\[\\d\\w]\+\$`

func somefunc(b bool) error {
	if b {
		return ErrSomeWrong
	} else {
		return SomeWrongError
	}
}

func main() {
	if err := somefunc(true); err != nil {
		log.Fatal(err)
	}
}
