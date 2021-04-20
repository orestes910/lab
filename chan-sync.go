package main

import (
	"fmt"
	"time"
)

// This function accepts a channel with
// a bool type called "done"
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	// the channel is sent "true"
	done <- true
}

func main() {

	// channel with a bool type, length 1
	done := make(chan bool, 1)
	// start a goroutine with the worker function.
	// the channel will be populated with "true" at
	// the end of the function execution
	go worker(done)

	// receive a message from the done channel.
	// This will block main() from finishing until
	// a message has been received, meaning the
	// goroutine must finish before the program
	// exits
	<-done
}
