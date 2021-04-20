package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// values can still be retrieved
	// from a closed channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
