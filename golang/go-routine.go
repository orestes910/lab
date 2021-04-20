package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// synchronous
	f("direct")

	// asynch
	go f("goroutine")

	// anonymous function invocation
	// with definition
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

}
