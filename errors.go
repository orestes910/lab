package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// return the int -1 and the error "no 42"
		return -1, errors.New("No 42")
	}
	// return sum along with a nil error value
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// arg e is type argError, func name is Error, returns string
func (e *argError) Error() string {
	// %d is int, %s is string
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// return int -1 and an argError type pointer
		return -1, &argError{arg, "No"}
	}
	return arg + 3, nil
}

func main() {
	// create a slice and range over it
	for _, i := range []int{7, 42} {
		// execute f1 against the current iteration
		// setting r to the return and e to the error
		// check if there is an error or not
		if r, e := f1(i); e != nil {
			// if there's an error, print it
			fmt.Println("f1 failed:", e)
		} else {
			// If not, print the return
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		// same with f2
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// e returns "No"
	_, e := f2(42)
	// assign ae and ok to the pointer argError
	// type in e. check for a value in ok, which
	// will be the prob val. There will be an
	// error because 42 produces one.
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
