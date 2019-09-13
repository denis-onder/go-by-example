package main

import (
	"errors"
	"fmt"
)

func basicErrors(arg1 int) (int, error) {
	if arg1 == 0 {
		return 0, errors.New("Provide a number other than 0")
	}
	return arg1, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func customErrors(arg1 int) (int, error) {
	if arg1 == 5 {
		return arg1, &argError{arg1, "Please provide a number other than 5"}
	}
	return arg1, nil
}

func main() {
	if num, err := basicErrors(10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
	if num, err := customErrors(5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
