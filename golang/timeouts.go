package main

import (
	"fmt"
	"time"
)

func main() {

	// adding any length, even 1, makes
	// the channel a buffered one
	c1 := make(chan string, 1)
	go func() {
		// sleep 2 seconds and send
		// "result 1" to c1 channel
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	// print either "result 1" from
	// c1, or "timeout 1", whichever happens
	// first. ("timeout 1" happens first)
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	// 3 seconds, so the timeout wont
	// be reached
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
