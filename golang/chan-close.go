package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// listen on jobs channel. If more = false,
			// the channel is empty and has been closed
			j, more := <-jobs
			// print the job if there is one
			if more {
				fmt.Println("received job", j)
				// if the channel is empty and closed, update
				// the done channel with true and return
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to the jobs channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(time.Second)
	}
	// close the channel, which will make the
	// "more" value above return false
	close(jobs)
	fmt.Println("sent all jobs")

	// wait for the worker go routine
	<-done
}
