package main

import "fmt"

func main() {

	// create a multi value buffered channel
	// by including a length in the def
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// calling print twice retrieves the values
	// in the order they were inserted, which
	// seems unintuitive and clunky
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
