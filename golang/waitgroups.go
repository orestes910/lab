package main

import (
	"fmt"
	"sync"
	"time"
)

// this function accepts a pointer to a waitgroup
func worker(id int, wg *sync.WaitGroup) {

	// When the function is finished, call the
	// Done method on the waitgroup. This will
	// add one negative delta per call.
	defer wg.Done()

	fmt.Println("Worker ", id, "starting")
	time.Sleep(time.Second)
	fmt.Println("Worker ", id, "done")
}

func main() {

	// declare the waitgroup
	var wg sync.WaitGroup

	// add single delta to the waitgroup
	// and launch a worker goroutine.
	// 5 times each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// The Wait method will execute until the
	// waitgroup counter is at 0
	wg.Wait()

	// Each wg.Add puts a record into the waitgroup, while
	// each wg.Done removes one. A record is added before
	// each goroutine is started, while the function being called
	// includes a call to remove a record.
}
