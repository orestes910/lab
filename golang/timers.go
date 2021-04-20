package main

import (
	"fmt"
	"time"
)

func main() {
	// a 2 second timer
	timer1 := time.NewTimer(2 * time.Second)

	// a timer creates a channel called "c"
	// block on that channel, meaning wait until
	// the timer "fires."
	<-timer1.C
	fmt.Println("Timer 1 has fired")

	// a one second timer
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 has fired")
	}()
	// This method kills the timer. Because
	// we're running this immediately after
	// starting it, the timer in the goroutine
	// won't fire.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 has been stopped")
	}

	fmt.Println("I'll sleep now")
	// Simple sleep
	time.Sleep(2 * time.Second)
	fmt.Println("I have slept")
}
