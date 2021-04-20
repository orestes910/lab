package main

import (
	"fmt"
	"time"
)

func main() {
	// a ticker that fires every half second
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// This function will either receive ticks until
	// true is sent to the "done" channel
	go func() {
		for {
			select {
			// Either receive a message
			// on the "done" channel
			case <-done:
				fmt.Println("Ticker stopping")
				return
			// Or receive a message on
			// the ticker's channel
			case t := <-ticker.C:
				fmt.Println("Ticker ticked at", t)
			}
		}
	}()
	// Allow a few ticks to occur
	time.Sleep(2 * time.Second)
	// Stop the ticker
	ticker.Stop()
	// Send true to the "done" channel,
	// killing the goroutine
	done <- true
	fmt.Println("Ticker stopped")
}
