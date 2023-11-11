package b

import (
	"errors"
	"log"
)

var ErrSomeWrong = errors.New("somethis is wrong") // want `Error variable ErrSomeWrong does not follow naming conventions, rule is \^\[\\d\\w]\+Error\$`
var SomeWrongError = errors.New("somethis is wrong")

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
