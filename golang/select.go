package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	// this syntax is to define and execute
	// a function all at once:
	// func(params) returns/types {code} (input)
	// func(input string) output string {
	//	  fmt.Println("Hello: ", input)
	// } ("Tim")
	go func() {
		// a verbose one second
		time.Sleep(1 * time.Second)
		// after a one second sleep,
		// write "one" to c1 channel
		c1 <- "one"
	}()

	go func() {
		// two seconds
		time.Sleep(2 * time.Second)

		c2 <- "two"
	}()

	// run select twice, so that when
	// the first message is received,
	// another select is still waiting
	// for the second message
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
