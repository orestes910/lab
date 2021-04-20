package main

import "fmt"

func main() {

	// make(chan type) to create a channel
	messages := make(chan string)

	// send "ping" into messages channel with <-
	go func() { messages <- "ping" }()

	// receive a message from a channel with
	// <-channel-name
	msg := <-messages
	fmt.Println(msg)

}
